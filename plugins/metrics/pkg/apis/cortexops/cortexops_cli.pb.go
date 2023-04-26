// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/cortexops/cortexops.proto

package cortexops

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/storage/v1"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	v2 "github.com/thediveo/enumflag/v2"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	time "time"
)

type contextKey_CortexOps_type struct{}

var contextKey_CortexOps contextKey_CortexOps_type

func ContextWithCortexOpsClient(ctx context.Context, client CortexOpsClient) context.Context {
	return context.WithValue(ctx, contextKey_CortexOps, client)
}

func CortexOpsClientFromContext(ctx context.Context) (CortexOpsClient, bool) {
	client, ok := ctx.Value(contextKey_CortexOps).(CortexOpsClient)
	return client, ok
}

func BuildCortexOpsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ops",
		Short:             `The CortexOps service contains setup and configuration lifecycle actions for the managed Cortex cluster.`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(BuildGetClusterConfigurationCmd())
	cmd.AddCommand(BuildConfigureClusterCmd())
	cmd.AddCommand(BuildGetClusterStatusCmd())
	cmd.AddCommand(BuildUninstallClusterCmd())
	return cmd
}

func BuildGetClusterConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-configuration",
		Short: "Gets the current configuration of the managed Cortex cluster.",
		Long: `
HTTP handlers for this method:
- get:"/configuration"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetClusterConfiguration(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cmd.Println(protojson.MarshalOptions{Multiline: true, EmitUnpopulated: true}.Format(response))
			return nil
		},
	}
	return cmd
}

func BuildConfigureClusterCmd() *cobra.Command {
	input := &ClusterConfiguration{}
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Updates the configuration of the managed Cortex cluster to match the provided configuration.",
		Long: `
If the cluster is not installed, it will be configured and installed.
Otherwise, the already-installed cluster will be reconfigured.

Note: some fields may contain secrets. The placeholder value "***" can be used to
keep an existing secret when updating the cluster configuration.

HTTP handlers for this method:
- post:"/configure"  body:"*"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.ConfigureCluster(cmd.Context(), input)
			return err
		},
	}
	cmd.Flags().AddFlagSet(input.FlagSet())
	return cmd
}

func (input *ClusterConfiguration) FlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("ClusterConfiguration", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(v2.New(&input.Mode, "DeploymentMode", map[DeploymentMode][]string{
		DeploymentMode_AllInOne:        {"AllInOne"},
		DeploymentMode_HighlyAvailable: {"HighlyAvailable"},
	}, v2.EnumCaseSensitive), "mode", "The deployment mode to use for Cortex.")
	if input.Storage == nil {
		input.Storage = &v1.StorageSpec{}
	}
	fs.StringVar(&input.Storage.Backend, "storage.backend", "", "")
	if input.Storage.S3 == nil {
		input.Storage.S3 = &v1.S3StorageSpec{}
	}
	fs.StringVar(&input.Storage.S3.Endpoint, "storage.s3.endpoint", "", "The S3 bucket endpoint. It could be an AWS S3 endpoint listed at  https:docs.aws.amazon.com/general/latest/gr/s3.html or the address of an  S3-compatible service in hostname:port format.")
	fs.StringVar(&input.Storage.S3.Region, "storage.s3.region", "", "S3 region. If unset, the client will issue a S3 GetBucketLocation API call  to autodetect it.")
	fs.StringVar(&input.Storage.S3.BucketName, "storage.s3.bucket-name", "", "S3 bucket name")
	fs.StringVar(&input.Storage.S3.SecretAccessKey, "storage.s3.secret-access-key", "", "S3 secret access key")
	fs.StringVar(&input.Storage.S3.AccessKeyID, "storage.s3.access-key-id", "", "S3 access key ID")
	fs.BoolVar(&input.Storage.S3.Insecure, "storage.s3.insecure", false, "If enabled, use http: for the S3 endpoint instead of https:. This could  be useful in local dev/test environments while using an S3-compatible  backend storage, like Minio.")
	fs.StringVar(&input.Storage.S3.SignatureVersion, "storage.s3.signature-version", "", "The signature version to use for authenticating against S3.  Supported values are: v4, v2")
	if input.Storage.S3.Sse == nil {
		input.Storage.S3.Sse = &v1.SSEConfig{}
	}
	fs.StringVar(&input.Storage.S3.Sse.Type, "storage.s3.sse.type", "", "Enable AWS Server Side Encryption. Supported values: SSE-KMS, SSE-S3")
	fs.StringVar(&input.Storage.S3.Sse.KmsKeyID, "storage.s3.sse.kms-key-id", "", "KMS Key ID used to encrypt objects in S3")
	fs.StringVar(&input.Storage.S3.Sse.KmsEncryptionContext, "storage.s3.sse.kms-encryption-context", "", "KMS Encryption Context used for object encryption. It expects a JSON formatted string.")
	if input.Storage.S3.Http == nil {
		input.Storage.S3.Http = &v1.HTTPConfig{}
	}
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.S3.Http.IdleConnTimeout), "storage.s3.http.idle-conn-timeout", "The time an idle connection will remain idle before closing.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.S3.Http.ResponseHeaderTimeout), "storage.s3.http.response-header-timeout", "The amount of time the client will wait for a servers response headers.")
	fs.BoolVar(&input.Storage.S3.Http.InsecureSkipVerify, "storage.s3.http.insecure-skip-verify", false, "If the client connects via HTTPS and this option is enabled, the client will accept any certificate and hostname.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.S3.Http.TlsHandshakeTimeout), "storage.s3.http.tls-handshake-timeout", "Maximum time to wait for a TLS handshake. 0 means no limit.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.S3.Http.ExpectContinueTimeout), "storage.s3.http.expect-continue-timeout", "The time to wait for a server's first response headers after fully writing the request headers if the request has an Expect header. 0 to send the request body immediately.")
	fs.Int32Var(&input.Storage.S3.Http.MaxIdleConns, "storage.s3.http.max-idle-conns", 0, "Maximum number of idle (keep-alive) connections across all hosts. 0 means no limit.")
	fs.Int32Var(&input.Storage.S3.Http.MaxIdleConnsPerHost, "storage.s3.http.max-idle-conns-per-host", 0, "Maximum number of idle (keep-alive) connections to keep per-host. If 0, a built-in default value is used.")
	fs.Int32Var(&input.Storage.S3.Http.MaxConnsPerHost, "storage.s3.http.max-conns-per-host", 0, "Maximum number of connections per host. 0 means no limit.")
	if input.Storage.Gcs == nil {
		input.Storage.Gcs = &v1.GCSStorageSpec{}
	}
	fs.StringVar(&input.Storage.Gcs.BucketName, "storage.gcs.bucket-name", "", "GCS bucket name")
	fs.StringVar(&input.Storage.Gcs.ServiceAccount, "storage.gcs.service-account", "", "JSON representing either a Google Developers Console client_credentials.json file  or a Google Developers service account key file. If empty, fallback to Google default logic.")
	if input.Storage.Azure == nil {
		input.Storage.Azure = &v1.AzureStorageSpec{}
	}
	fs.StringVar(&input.Storage.Azure.StorageAccountName, "storage.azure.storage-account-name", "", "Azure storage account name")
	fs.StringVar(&input.Storage.Azure.StorageAccountKey, "storage.azure.storage-account-key", "", "Azure storage account key")
	fs.StringVar(&input.Storage.Azure.ContainerName, "storage.azure.container-name", "", "Azure storage container name")
	fs.StringVar(&input.Storage.Azure.Endpoint, "storage.azure.endpoint", "", "Azure storage endpoint suffix without schema. The account name will be  prefixed to this value to create the FQDN")
	fs.Int32Var(&input.Storage.Azure.MaxRetries, "storage.azure.max-retries", 0, "Number of retries for recoverable errors")
	fs.StringVar(&input.Storage.Azure.MsiResource, "storage.azure.msi-resource", "", "Azure storage MSI resource. Either this or account key must be set.")
	fs.StringVar(&input.Storage.Azure.UserAssignedID, "storage.azure.user-assigned-id", "", "Azure storage MSI resource managed identity client Id. If not supplied system assigned identity is used")
	if input.Storage.Azure.Http == nil {
		input.Storage.Azure.Http = &v1.HTTPConfig{}
	}
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Azure.Http.IdleConnTimeout), "storage.azure.http.idle-conn-timeout", "The time an idle connection will remain idle before closing.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Azure.Http.ResponseHeaderTimeout), "storage.azure.http.response-header-timeout", "The amount of time the client will wait for a servers response headers.")
	fs.BoolVar(&input.Storage.Azure.Http.InsecureSkipVerify, "storage.azure.http.insecure-skip-verify", false, "If the client connects via HTTPS and this option is enabled, the client will accept any certificate and hostname.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Azure.Http.TlsHandshakeTimeout), "storage.azure.http.tls-handshake-timeout", "Maximum time to wait for a TLS handshake. 0 means no limit.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Azure.Http.ExpectContinueTimeout), "storage.azure.http.expect-continue-timeout", "The time to wait for a server's first response headers after fully writing the request headers if the request has an Expect header. 0 to send the request body immediately.")
	fs.Int32Var(&input.Storage.Azure.Http.MaxIdleConns, "storage.azure.http.max-idle-conns", 0, "Maximum number of idle (keep-alive) connections across all hosts. 0 means no limit.")
	fs.Int32Var(&input.Storage.Azure.Http.MaxIdleConnsPerHost, "storage.azure.http.max-idle-conns-per-host", 0, "Maximum number of idle (keep-alive) connections to keep per-host. If 0, a built-in default value is used.")
	fs.Int32Var(&input.Storage.Azure.Http.MaxConnsPerHost, "storage.azure.http.max-conns-per-host", 0, "Maximum number of connections per host. 0 means no limit.")
	if input.Storage.Swift == nil {
		input.Storage.Swift = &v1.SwiftStorageSpec{}
	}
	fs.Int32Var(&input.Storage.Swift.AuthVersion, "storage.swift.auth-version", 0, "OpenStack Swift authentication API version. 0 to autodetect.")
	fs.StringVar(&input.Storage.Swift.AuthURL, "storage.swift.auth-url", "", "OpenStack Swift authentication URL.")
	fs.StringVar(&input.Storage.Swift.Username, "storage.swift.username", "", "OpenStack Swift username.")
	fs.StringVar(&input.Storage.Swift.UserDomainName, "storage.swift.user-domain-name", "", "OpenStack Swift user's domain name.")
	fs.StringVar(&input.Storage.Swift.UserDomainID, "storage.swift.user-domain-id", "", "OpenStack Swift user's domain ID.")
	fs.StringVar(&input.Storage.Swift.UserID, "storage.swift.user-id", "", "OpenStack Swift user ID.")
	fs.StringVar(&input.Storage.Swift.Password, "storage.swift.password", "", "OpenStack Swift API key.")
	fs.StringVar(&input.Storage.Swift.DomainID, "storage.swift.domain-id", "", "OpenStack Swift user's domain ID.")
	fs.StringVar(&input.Storage.Swift.DomainName, "storage.swift.domain-name", "", "OpenStack Swift user's domain name.")
	fs.StringVar(&input.Storage.Swift.ProjectID, "storage.swift.project-id", "", "OpenStack Swift project ID (v2,v3 auth only).")
	fs.StringVar(&input.Storage.Swift.ProjectName, "storage.swift.project-name", "", "OpenStack Swift project name (v2,v3 auth only).")
	fs.StringVar(&input.Storage.Swift.ProjectDomainID, "storage.swift.project-domain-id", "", "ID of the OpenStack Swift project's domain (v3 auth only), only needed  if it differs the from user domain.")
	fs.StringVar(&input.Storage.Swift.ProjectDomainName, "storage.swift.project-domain-name", "", "Name of the OpenStack Swift project's domain (v3 auth only), only needed  if it differs from the user domain.")
	fs.StringVar(&input.Storage.Swift.RegionName, "storage.swift.region-name", "", "OpenStack Swift Region to use (v2,v3 auth only).")
	fs.StringVar(&input.Storage.Swift.ContainerName, "storage.swift.container-name", "", "Name of the OpenStack Swift container to use. The container must already  exist.")
	fs.Int32Var(&input.Storage.Swift.MaxRetries, "storage.swift.max-retries", 0, "Max number of times to retry failed requests.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Swift.ConnectTimeout), "storage.swift.connect-timeout", "Time after which a connection attempt is aborted.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.Swift.RequestTimeout), "storage.swift.request-timeout", "Time after which an idle request is aborted. The timeout watchdog is reset  each time some data is received, so the timeout triggers after X time no  data is received on a request.")
	if input.Storage.Filesystem == nil {
		input.Storage.Filesystem = &v1.FilesystemStorageSpec{}
	}
	fs.StringVar(&input.Storage.Filesystem.Directory, "storage.filesystem.directory", "", "Local filesystem storage directory.")
	fs.Var(flagutil.DurationpbValue(0, &input.Storage.RetentionPeriod), "storage.retention-period", "")
	if input.Grafana == nil {
		input.Grafana = &GrafanaConfig{}
	}
	fs.BoolVar(&input.Grafana.Enabled, "grafana.enabled", false, "")
	fs.StringVar(&input.Grafana.Hostname, "grafana.hostname", "", "DNS name at which Grafana will be available in the browser.")
	if input.Workloads == nil {
		input.Workloads = &Workloads{}
	}
	if input.Workloads.Distributor == nil {
		input.Workloads.Distributor = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Distributor.Replicas, "workloads.distributor.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Distributor.ExtraArgs, "workloads.distributor.extra-args", nil, "")
	if input.Workloads.Ingester == nil {
		input.Workloads.Ingester = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Ingester.Replicas, "workloads.ingester.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Ingester.ExtraArgs, "workloads.ingester.extra-args", nil, "")
	if input.Workloads.Compactor == nil {
		input.Workloads.Compactor = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Compactor.Replicas, "workloads.compactor.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Compactor.ExtraArgs, "workloads.compactor.extra-args", nil, "")
	if input.Workloads.StoreGateway == nil {
		input.Workloads.StoreGateway = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.StoreGateway.Replicas, "workloads.store-gateway.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.StoreGateway.ExtraArgs, "workloads.store-gateway.extra-args", nil, "")
	if input.Workloads.Ruler == nil {
		input.Workloads.Ruler = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Ruler.Replicas, "workloads.ruler.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Ruler.ExtraArgs, "workloads.ruler.extra-args", nil, "")
	if input.Workloads.QueryFrontend == nil {
		input.Workloads.QueryFrontend = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.QueryFrontend.Replicas, "workloads.query-frontend.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.QueryFrontend.ExtraArgs, "workloads.query-frontend.extra-args", nil, "")
	if input.Workloads.Querier == nil {
		input.Workloads.Querier = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Querier.Replicas, "workloads.querier.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Querier.ExtraArgs, "workloads.querier.extra-args", nil, "")
	if input.Workloads.Purger == nil {
		input.Workloads.Purger = &CortexWorkloadSpec{}
	}
	fs.Int32Var(&input.Workloads.Purger.Replicas, "workloads.purger.replicas", 0, "")
	fs.StringSliceVar(&input.Workloads.Purger.ExtraArgs, "workloads.purger.extra-args", nil, "")
	if input.Cortex == nil {
		input.Cortex = &CortexConfig{}
	}
	if input.Cortex.Compactor == nil {
		input.Cortex.Compactor = &CompactorConfig{}
	}
	fs.Var(flagutil.DurationpbSliceValue([]time.Duration{2 * time.Hour, 12 * time.Hour, 24 * time.Hour}, &input.Cortex.Compactor.BlockRanges), "cortex.compactor.block-ranges", "List of compaction time ranges")
	fs.Var(flagutil.DurationpbValue(1*time.Hour, &input.Cortex.Compactor.CompactionInterval), "cortex.compactor.compaction-interval", "The frequency at which the compaction runs")
	fs.Var(flagutil.DurationpbValue(15*time.Minute, &input.Cortex.Compactor.CleanupInterval), "cortex.compactor.cleanup-interval", "How frequently compactor should run blocks cleanup and maintenance, as well as update the bucket index")
	fs.Var(flagutil.DurationpbValue(12*time.Hour, &input.Cortex.Compactor.DeletionDelay), "cortex.compactor.deletion-delay", "Time before a block marked for deletion is deleted from the bucket")
	fs.Var(flagutil.DurationpbValue(6*time.Hour, &input.Cortex.Compactor.TenantCleanupDelay), "cortex.compactor.tenant-cleanup-delay", "For tenants marked for deletion, this is time between deleting of last block, and doing final cleanup (marker files, debug files) of the tenant")
	if input.Cortex.Querier == nil {
		input.Cortex.Querier = &QuerierConfig{}
	}
	fs.Var(flagutil.DurationpbValue(2*time.Minute, &input.Cortex.Querier.QueryTimeout), "cortex.querier.query-timeout", "The timeout for a query")
	fs.Int32Var(&input.Cortex.Querier.MaxSamples, "cortex.querier.max-samples", 0, "Maximum number of samples a single query can load into memory")
	fs.Var(flagutil.DurationpbValue(0, &input.Cortex.Querier.QueryIngestersWithin), "cortex.querier.query-ingesters-within", "Maximum lookback beyond which queries are not sent to ingester. 0 means all queries are sent to ingester.")
	fs.Var(flagutil.DurationpbValue(10*time.Minute, &input.Cortex.Querier.MaxQueryIntoFuture), "cortex.querier.max-query-into-future", "Maximum duration into the future you can query. 0 to disable")
	fs.Var(flagutil.DurationpbValue(1*time.Minute, &input.Cortex.Querier.DefaultEvaluationInterval), "cortex.querier.default-evaluation-interval", "The default evaluation interval or step size for subqueries")
	fs.Var(flagutil.DurationpbValue(0, &input.Cortex.Querier.QueryStoreAfter), "cortex.querier.query-store-after", "The time after which a metric should be queried from storage and not just ingesters. 0 means all queries are sent to store.  When running the blocks storage, if this option is enabled, the time range of the query sent to the store will be manipulated  to ensure the query end is not more recent than 'now - query-store-after'.")
	fs.Var(flagutil.DurationpbValue(5*time.Minute, &input.Cortex.Querier.LookbackDelta), "cortex.querier.lookback-delta", "Time since the last sample after which a time series is considered stale and ignored by expression evaluations")
	fs.Var(flagutil.DurationpbValue(0, &input.Cortex.Querier.ShuffleShardingIngestersLookbackPeriod), "cortex.querier.shuffle-sharding-ingesters-lookback-period", "When distributor's sharding strategy is shuffle-sharding and this setting is > 0, queriers fetch in-memory series from  the minimum set of required ingesters, selecting only ingesters which may have received series since 'now - lookback period'.  The lookback period should be greater or equal than the configured 'query store after' and 'query ingesters within'.  If this setting is 0, queriers always query all ingesters (ingesters shuffle sharding on read path is disabled).")
	fs.Int32Var(&input.Cortex.Querier.MaxFetchedSeriesPerQuery, "cortex.querier.max-fetched-series-per-query", 0, "The maximum number of unique series for which a query can fetch samples from each ingesters and blocks storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable")
	if input.Cortex.Distributor == nil {
		input.Cortex.Distributor = &DistributorConfig{}
	}
	fs.Float64Var(&input.Cortex.Distributor.IngestionRate, "cortex.distributor.ingestion-rate", 0.0, "Per-user ingestion rate limit in samples per second.")
	fs.StringVar(&input.Cortex.Distributor.IngestionRateStrategy, "cortex.distributor.ingestion-rate-strategy", "", "Whether the ingestion rate limit should be applied individually to each distributor instance (local), or evenly shared across the cluster (global).")
	fs.Int32Var(&input.Cortex.Distributor.IngestionBurstSize, "cortex.distributor.ingestion-burst-size", 0, "Per-user allowed ingestion burst size (in number of samples).")
	if input.Cortex.Ingester == nil {
		input.Cortex.Ingester = &IngesterConfig{}
	}
	fs.Int32Var(&input.Cortex.Ingester.MaxLocalSeriesPerUser, "cortex.ingester.max-local-series-per-user", 0, "The maximum number of active series per user, per ingester. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxLocalSeriesPerMetric, "cortex.ingester.max-local-series-per-metric", 0, "The maximum number of active series per metric name, per ingester. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxGlobalSeriesPerUser, "cortex.ingester.max-global-series-per-user", 0, "The maximum number of active series per user, across the cluster before replication. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxGlobalSeriesPerMetric, "cortex.ingester.max-global-series-per-metric", 0, "The maximum number of active series per metric name, across the cluster before replication. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxLocalMetricsWithMetadataPerUser, "cortex.ingester.max-local-metrics-with-metadata-per-user", 0, "The maximum number of active metrics with metadata per user, per ingester. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxLocalMetadataPerMetric, "cortex.ingester.max-local-metadata-per-metric", 0, "The maximum number of metadata per metric, per ingester. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxGlobalMetricsWithMetadataPerUser, "cortex.ingester.max-global-metrics-with-metadata-per-user", 0, "The maximum number of active metrics with metadata per user, across the cluster. 0 to disable.")
	fs.Int32Var(&input.Cortex.Ingester.MaxGlobalMetadataPerMetric, "cortex.ingester.max-global-metadata-per-metric", 0, "The maximum number of metadata per metric, across the cluster. 0 to disable.")
	return fs
}

func BuildGetClusterStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Gets the current status of the managed Cortex cluster.",
		Long: `
The status includes the current install state, version, and metadata. If
the cluster is in the process of being reconfigured or uninstalled, it will
be reflected in the install state.
No guarantees are made about the contents of the metadata field; its
contents are strictly informational.

HTTP handlers for this method:
- get:"/status"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.GetClusterStatus(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cmd.Println(protojson.MarshalOptions{Multiline: true, EmitUnpopulated: true}.Format(response))
			return nil
		},
	}
	return cmd
}

func BuildUninstallClusterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstalls the managed Cortex cluster.",
		Long: `
Implementation details including error handling and system state requirements
are left to the cluster driver, and this API makes no guarantees about
the state of the cluster after the call completes (regardless of success).

HTTP handlers for this method:
- post:"/uninstall"
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.UninstallCluster(cmd.Context(), &emptypb.Empty{})
			return err
		},
	}
	return cmd
}
