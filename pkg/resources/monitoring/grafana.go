package monitoring

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	_ "embed"

	grafanav1beta1 "github.com/grafana-operator/grafana-operator/v5/api/v1beta1"
	"github.com/imdario/mergo"
	"github.com/rancher/opni/pkg/resources"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

//go:embed dashboards/dashboards.json
var dashboardsJson []byte

//go:embed dashboards/opni-gateway.json
var opniGatewayJson []byte

//go:embed dashboards/home.json
var homeDashboardJson []byte

//go:embed dashboards/opni-service-latency.json
var serviceLatencyDashboardJson []byte

//go:embed slo/slo_grafana_overview.json
var sloOverviewDashboard []byte

//go:embed slo/slo_grafana_detailed.json
var sloDetailedDashboard []byte

const (
	grafanaImageRepo    = "grafana"
	grafanaImageVersion = "10.1.5"
	secret              = "opni-gateway-client-cert"
)

func (r *Reconciler) grafana() ([]resources.Resource, error) {
	dashboardSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{
			resources.AppNameLabel:  "grafana",
			resources.PartOfLabel:   "opni",
			resources.InstanceLabel: r.mc.Name,
		},
	}

	grafana := &grafanav1beta1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "grafana",
			Namespace: r.mc.Namespace,
			Labels:    dashboardSelector.MatchLabels,
		},
	}

	legacyResources := []resources.Resource{
		resources.Absent(&grafanav1beta1.Grafana{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "opni-monitoring",
				Namespace: r.mc.Namespace,
			},
		}),
		resources.Absent(&grafanav1beta1.GrafanaDatasource{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "opni-monitoring",
				Namespace: r.mc.Namespace,
			},
		}),
	}

	opniDatasource, err := r.createOpniDatasource(dashboardSelector)
	if err != nil {
		return nil, err
	}

	opniAlertManagerDatasource, err := r.createOpniAlertManagerDatasource(dashboardSelector)
	if err != nil {
		return nil, err
	}

	grafanaDashboards, err := r.createGrafanaDashboards(dashboardSelector)
	if err != nil {
		return nil, err
	}

	if !r.mc.Spec.Grafana.GetEnabled() {
		absentResources := append([]resources.Resource{
			resources.Absent(grafana),
			resources.Absent(opniDatasource),
			resources.Absent(opniAlertManagerDatasource),
		}, legacyResources...)
		for _, dashboard := range grafanaDashboards {
			absentResources = append(absentResources, resources.Absent(dashboard))
		}
		return absentResources, nil
	}

	defaults := r.createGrafanaSpecDefaults()
	spec := r.mc.Spec.Grafana.GrafanaSpec

	// TODO: endpoint changes cause the deployment to be rolled out every time
	// ips := []string{}
	// for _, endpoint := range r.gw.Status.Endpoints {
	// 	ips = append(ips, endpoint.IP)
	// }
	// defaults.Config["auth.proxy"]["whitelist"] = strings.Join(ips, ",")

	// apply defaults to user-provided config
	// ensure label selectors and secrets are appended to any user defined ones
	if err := mergo.Merge(&spec, defaults, mergo.WithAppendSlice); err != nil {
		return nil, err
	}

	// special case as we don't want the append slice logic for access modes
	if spec.PersistentVolumeClaim.Spec.AccessModes == nil {
		spec.PersistentVolumeClaim.Spec.AccessModes = []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce}
	}

	grafana.Spec = spec

	// switch gatewayAuthProvider {
	// case v1beta1.AuthProviderNoAuth:
	// 	grafanaAuthGenericOauthCfg := map[string]string{
	// 		"enabled":             "true",
	// 		"client_id":           "grafana",
	// 		"client_secret":       "noauth",
	// 		"scopes":              "openid profile email",
	// 		"auth_url":            fmt.Sprintf("http://%s:4000/oauth2/authorize", gatewayHostname),
	// 		"token_url":           fmt.Sprintf("http://%s:4000/oauth2/token", gatewayHostname),
	// 		"api_url":             fmt.Sprintf("http://%s:4000/oauth2/userinfo", gatewayHostname),
	// 		"role_attribute_path": "grafana_role",
	// 	}
	// 	grafana.Spec.Config["auth.generic_oauth"] = grafanaAuthGenericOauthCfg

	// case v1beta1.AuthProviderOpenID:
	// 	spec := r.gw.Spec.Auth.Openid
	// 	if spec.Discovery == nil && spec.WellKnownConfiguration == nil {
	// 		return nil, openid.ErrMissingDiscoveryConfig
	// 	}
	// 	wkc, err := spec.OpenidConfig.GetWellKnownConfiguration()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to fetch configuration from openid provider: %w", err)
	// 	}
	// 	scopes := spec.Scopes
	// 	if len(scopes) == 0 {
	// 		scopes = []string{"openid", "profile", "email"}
	// 	}
	// 	grafanaAuthGenericOauthCfg := map[string]string{
	// 		"enabled":             "true",
	// 		"client_id":           spec.ClientID,
	// 		"client_secret":       spec.ClientSecret,
	// 		"scopes":              strings.Join(scopes, " "),
	// 		"auth_url":            wkc.AuthEndpoint,
	// 		"token_url":           wkc.TokenEndpoint,
	// 		"api_url":             wkc.UserinfoEndpoint,
	// 		"role_attribute_path": spec.RoleAttributePath,
	// 	}
	// 	if len(spec.AllowedDomains) > 0 {
	// 		grafanaAuthGenericOauthCfg["allowed_domains"] = strings.Join(spec.AllowedDomains, " ")
	// 	}
	// 	if spec.AllowSignUp != nil {
	// 		grafanaAuthGenericOauthCfg["allow_sign_up"] = strconv.FormatBool(lo.FromPtr(spec.AllowSignUp))
	// 	}
	// 	if spec.InsecureSkipVerify != nil {
	// 		grafanaAuthGenericOauthCfg["tls_skip_verify_insecure"] = strconv.FormatBool(lo.FromPtr(spec.InsecureSkipVerify))
	// 	}
	// 	if spec.RoleAttributeStrict != nil {
	// 		grafanaAuthGenericOauthCfg["role_attribute_strict"] = strconv.FormatBool(lo.FromPtr(spec.RoleAttributeStrict))
	// 	}
	// 	if spec.TLSClientCA != "" && spec.TLSClientCert != "" && spec.TLSClientKey != "" {
	// 		grafanaAuthGenericOauthCfg["tls_client_cert"] = spec.TLSClientCert
	// 		grafanaAuthGenericOauthCfg["tls_client_key"] = spec.TLSClientKey
	// 		grafanaAuthGenericOauthCfg["tls_client_ca"] = spec.TLSClientCA
	// 	}
	// 	if spec.EmailAttributePath != "" {
	// 		grafanaAuthGenericOauthCfg["email_attribute_path"] = spec.EmailAttributePath
	// 	}

	// 	grafana.Spec.Config["auth.generic_oauth"] = grafanaAuthGenericOauthCfg

	// 	if wkc.EndSessionEndpoint != "" {
	// 		grafana.Spec.Config["auth"]["signout_redirect_url"] = wkc.EndSessionEndpoint
	// 	}

	// 	if spec.InsecureSkipVerify != nil && *spec.InsecureSkipVerify {
	// 		r.lg.Warn(chalk.Yellow.Color("InsecureSkipVerify enabled for openid auth"))
	// 	}
	// }

	controllerutil.SetOwnerReference(r.mc, grafana, r.client.Scheme())
	controllerutil.SetOwnerReference(r.mc, opniDatasource, r.client.Scheme())
	controllerutil.SetOwnerReference(r.mc, opniAlertManagerDatasource, r.client.Scheme())

	presentResources := []resources.Resource{
		resources.Present(grafana),
		resources.Present(opniDatasource),
		resources.Present(opniAlertManagerDatasource),
	}
	for _, dashboard := range grafanaDashboards {
		controllerutil.SetOwnerReference(r.mc, dashboard, r.client.Scheme())
		presentResources = append(presentResources, resources.Present(dashboard))
	}

	return append(presentResources, legacyResources...), nil
}

func (r *Reconciler) createGrafanaSpecDefaults() *grafanav1beta1.GrafanaSpec {
	tag := grafanaImageVersion
	if r.mc.Spec.Grafana.GetVersion() != "" {
		tag = strings.TrimSpace(r.mc.Spec.Grafana.GetVersion())
	}

	return &grafanav1beta1.GrafanaSpec{
		Client: &grafanav1beta1.GrafanaClient{
			PreferIngress: lo.ToPtr(false),
		},
		Config: createDefaultGrafanaIni(),
		Deployment: &grafanav1beta1.DeploymentV1{
			Spec: grafanav1beta1.DeploymentV1Spec{
				Template: &grafanav1beta1.DeploymentV1PodTemplateSpec{
					Spec: &grafanav1beta1.DeploymentV1PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "grafana",
								Image: grafanaImageRepo + "/grafana:" + tag,
								Env: []corev1.EnvVar{
									{
										Name:  "GF_INSTALL_PLUGINS",
										Value: "grafana-polystat-panel,marcusolsson-treemap-panel",
									},
								},
							},
						},
						SecurityContext: &corev1.PodSecurityContext{
							FSGroup: lo.ToPtr(int64(472)),
						},
					},
				},
			},
		},
		PersistentVolumeClaim: &grafanav1beta1.PersistentVolumeClaimV1{
			Spec: &grafanav1beta1.PersistentVolumeClaimV1Spec{
				Resources: &corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse("10Gi"),
					},
				},
			},
		},
		Preferences: &grafanav1beta1.GrafanaPreferences{
			HomeDashboardUID: "opni-home",
		},
	}
}

func (r *Reconciler) getGrafanaHostname(gatewayHostname string) (string, error) {
	grafanaHostname := fmt.Sprintf("grafana.%s", gatewayHostname)
	if r.mc.Spec.Grafana.GetHostname() != "" {
		grafanaHostname = r.mc.Spec.Grafana.GetHostname()
	}

	if strings.Contains(grafanaHostname, "://") {
		_, grafanaHostname, _ = strings.Cut(grafanaHostname, "://")
	}

	grafanaHostname = strings.TrimSpace(grafanaHostname)
	if _, err := url.Parse(grafanaHostname); err != nil {
		return "", fmt.Errorf("invalid grafana hostname: %w", err)
	}

	return grafanaHostname, nil
}

func (r *Reconciler) createOpniDatasource(dashboardSelector *metav1.LabelSelector) (*grafanav1beta1.GrafanaDatasource, error) {
	opniDatasourceJSONCfg, err := createDatasourceJSONData()
	if err != nil {
		return nil, err
	}
	opniDatasourceSecureJSONCfg, err := createDatasourceSecureJSONData()
	if err != nil {
		return nil, err
	}

	return &grafanav1beta1.GrafanaDatasource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "opni-datasource",
			Namespace: r.mc.Namespace,
		},
		Spec: grafanav1beta1.GrafanaDatasourceSpec{
			ValuesFrom: grafanaDatasourceTLSSecret(),
			InstanceSelector: &metav1.LabelSelector{
				MatchLabels: dashboardSelector.MatchLabels,
			},
			Datasource: &grafanav1beta1.GrafanaDatasourceInternal{
				Name:           "Opni",
				Type:           "prometheus",
				Access:         "proxy",
				URL:            fmt.Sprintf("https://opni-internal.%s.svc:8080/api/prom", r.mc.Namespace),
				Editable:       lo.ToPtr(false),
				IsDefault:      lo.ToPtr(true),
				JSONData:       opniDatasourceJSONCfg,
				SecureJSONData: opniDatasourceSecureJSONCfg,
			},
		},
	}, nil
}

func (r *Reconciler) createOpniAlertManagerDatasource(dashboardSelector *metav1.LabelSelector) (*grafanav1beta1.GrafanaDatasource, error) {
	opniAlertManagerDatasourceJSONCfg, err := createAlertManagerDatasourceJSONData()
	if err != nil {
		return nil, err
	}

	opniDatasourceSecureJSONCfg, err := createDatasourceSecureJSONData()
	if err != nil {
		return nil, err
	}

	return &grafanav1beta1.GrafanaDatasource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "opni-alert-manager-datasource",
			Namespace: r.mc.Namespace,
		},
		Spec: grafanav1beta1.GrafanaDatasourceSpec{
			InstanceSelector: &metav1.LabelSelector{
				MatchLabels: dashboardSelector.MatchLabels,
			},
			ValuesFrom: grafanaDatasourceTLSSecret(),
			Datasource: &grafanav1beta1.GrafanaDatasourceInternal{
				Name:           "Opni Alertmanager",
				UID:            "opni_alertmanager",
				Type:           "alertmanager",
				Access:         "proxy",
				URL:            fmt.Sprintf("https://opni-internal.%s.svc:8080/api/prom", r.mc.Namespace),
				Editable:       lo.ToPtr(false),
				JSONData:       opniAlertManagerDatasourceJSONCfg,
				SecureJSONData: opniDatasourceSecureJSONCfg,
			},
		},
	}, nil
}

func grafanaDatasourceTLSSecret() []grafanav1beta1.GrafanaDatasourceValueFrom {
	return []grafanav1beta1.GrafanaDatasourceValueFrom{
		{
			TargetPath: "secureJsonData.tlsCACert",
			ValueFrom: grafanav1beta1.GrafanaDatasourceValueFromSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secret,
					},
					Key: "ca.crt",
				},
			},
		},
		{
			TargetPath: "secureJsonData.tlsClientCert",
			ValueFrom: grafanav1beta1.GrafanaDatasourceValueFromSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secret,
					},
					Key: "tls.crt",
				},
			},
		},
		{
			TargetPath: "secureJsonData.tlsClientKey",
			ValueFrom: grafanav1beta1.GrafanaDatasourceValueFromSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secret,
					},
					Key: "tls.key",
				},
			},
		},
	}
}

// createDefaultGrafanaIni defines the grafana.ini file.
// See: https://grafana.com/docs/grafana/latest/setup-grafana/configure-grafana/
func createDefaultGrafanaIni() map[string]map[string]string {
	config := make(map[string]map[string]string)

	config["security"] = map[string]string{
		// TODO: grafana operator absolutely hammers the api when provisioning and will get itself locked out immediately if the server is not ready, or something like that
		"disable_brute_force_login_protection": "true",
	}
	config["dataproxy"] = map[string]string{
		"send_user_header": "true",
		"logging":          "true",
	}

	config["server"] = map[string]string{
		"root_url":            "/proxy/grafana",
		"serve_from_sub_path": "false", // very important
	}

	config["auth"] = map[string]string{
		"disable_login_form": "true",
	}

	config["users"] = map[string]string{
		"allow_sign_up":      "false",
		"auto_assign_org":    "true",
		"auto_assign_org_id": "1",
	}
	config["auth.proxy"] = map[string]string{
		"enabled":         "true",
		"header_name":     "X-WEBAUTH-USER",
		"header_property": "username",
		"auto_sign_up":    "true",
		"sync_ttl":        "60",
		"headers":         "Role:X-WEBAUTH-ROLE",

		// https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/auth-proxy/#login-token-and-session-cookie
		"enable_login_token": "false",
	}
	// oauthSection := make(map[string]string)
	// oauthSection["enabled"] = "true"
	// oauthSection["scopes"] = "openid profile email"
	// config["auth.generic_oauth"] = oauthSection

	unifiedAlertingSection := make(map[string]string)
	unifiedAlertingSection["enabled"] = "true"
	config["unified_alerting"] = unifiedAlertingSection

	alertingSection := make(map[string]string)
	alertingSection["enabled"] = "false"
	config["alerting"] = alertingSection

	featureTogglesSection := make(map[string]string)
	featureTogglesSection["enable"] = "panelTitleSearch increaseInMemDatabaseQueryCache newPanelChromeUI idForwarding"
	config["feature_toggles"] = featureTogglesSection

	dashboardsSection := make(map[string]string)
	dashboardsSection["versions_to_keep"] = "10"
	config["dashboards"] = dashboardsSection

	return config
}

func createDatasourceJSONData() (json.RawMessage, error) {
	datasourceCfg := make(map[string]any)
	datasourceCfg["alertmanagerUid"] = "opni_alertmanager"
	datasourceCfg["tlsAuth"] = true
	datasourceCfg["tlsAuthWithCACert"] = true
	datasourceCfg["forwardGrafanaIdToken"] = true

	jsonData, err := json.Marshal(datasourceCfg)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func createAlertManagerDatasourceJSONData() (json.RawMessage, error) {
	datasourceCfg := make(map[string]any)
	datasourceCfg["implementation"] = "cortex"
	datasourceCfg["withCredentials"] = true
	datasourceCfg["tlsAuth"] = true
	datasourceCfg["tlsAuthWithCACert"] = true

	jsonData, err := json.Marshal(datasourceCfg)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func createDatasourceSecureJSONData() (json.RawMessage, error) {
	datasourceSecureCfg := make(map[string]string)
	datasourceSecureCfg["tlsCACert"] = "${ca.crt}"
	datasourceSecureCfg["tlsClientCert"] = "${tls.crt}"
	datasourceSecureCfg["tlsClientKey"] = "${tls.key}"

	jsonData, err := json.Marshal(datasourceSecureCfg)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (r *Reconciler) createGrafanaDashboards(dashboardSelector *metav1.LabelSelector) ([]*grafanav1beta1.GrafanaDashboard, error) {
	opniPrometheusDatasource := grafanav1beta1.GrafanaDashboardDatasource{
		DatasourceName: "Opni",
		InputName:      "DS_PROMETHEUS",
	}

	grafanaDashboards := []*grafanav1beta1.GrafanaDashboard{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "opni-gateway.json",
				Namespace: r.mc.Namespace,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(opniGatewayJson),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
				Datasources: []grafanav1beta1.GrafanaDashboardDatasource{opniPrometheusDatasource},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "opni-home.json",
				Namespace: r.mc.Namespace,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(homeDashboardJson),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
				Datasources: []grafanav1beta1.GrafanaDashboardDatasource{opniPrometheusDatasource},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "opni-service-latency.json",
				Namespace: r.mc.Namespace,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(serviceLatencyDashboardJson),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
				Datasources: []grafanav1beta1.GrafanaDashboardDatasource{opniPrometheusDatasource},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "slo-overview.json",
				Namespace: r.mc.Namespace,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(sloOverviewDashboard),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
				Datasources: []grafanav1beta1.GrafanaDashboardDatasource{opniPrometheusDatasource},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "slo-detailed.json",
				Namespace: r.mc.Namespace,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(sloDetailedDashboard),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
				Datasources: []grafanav1beta1.GrafanaDashboardDatasource{opniPrometheusDatasource},
			},
		},
	}

	dashboards := map[string]json.RawMessage{}
	if err := json.Unmarshal(dashboardsJson, &dashboards); err != nil {
		return nil, err
	}
	for name, jsonData := range dashboards {
		grafanaDashboards = append(grafanaDashboards, &grafanav1beta1.GrafanaDashboard{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: r.mc.Namespace,
				Labels:    dashboardSelector.MatchLabels,
			},
			Spec: grafanav1beta1.GrafanaDashboardSpec{
				Json: string(jsonData),
				InstanceSelector: &metav1.LabelSelector{
					MatchLabels: dashboardSelector.MatchLabels,
				},
			},
		})
	}

	return grafanaDashboards, nil
}
