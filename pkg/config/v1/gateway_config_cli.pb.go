// Code generated by internal/codegen/cli/generator.go. DO NOT EDIT.
// source: github.com/rancher/opni/pkg/config/v1/gateway_config.proto

package v1

import (
	context "context"
	errors "errors"
	cli "github.com/rancher/opni/internal/codegen/cli"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	cliutil "github.com/rancher/opni/pkg/opni/cliutil"
	driverutil "github.com/rancher/opni/pkg/plugins/driverutil"
	storage "github.com/rancher/opni/pkg/storage"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	lo "github.com/samber/lo"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strings "strings"
)

type (
	contextKey_GatewayConfig_type      struct{}
	contextInjector_GatewayConfig_type struct{}
)

var (
	contextKey_GatewayConfig     contextKey_GatewayConfig_type
	GatewayConfigContextInjector contextInjector_GatewayConfig_type
)

func (contextInjector_GatewayConfig_type) NewClient(cc grpc.ClientConnInterface) GatewayConfigClient {
	return NewGatewayConfigClient(cc)
}

func (contextInjector_GatewayConfig_type) UnderlyingConn(client GatewayConfigClient) grpc.ClientConnInterface {
	return client.(*gatewayConfigClient).cc
}

func (contextInjector_GatewayConfig_type) ContextWithClient(ctx context.Context, client GatewayConfigClient) context.Context {
	return context.WithValue(ctx, contextKey_GatewayConfig, client)
}

func (contextInjector_GatewayConfig_type) ClientFromContext(ctx context.Context) (GatewayConfigClient, bool) {
	client, ok := ctx.Value(contextKey_GatewayConfig).(GatewayConfigClient)
	return client, ok
}

var extraCmds_GatewayConfig []*cobra.Command

func addExtraGatewayConfigCmd(custom *cobra.Command) {
	extraCmds_GatewayConfig = append(extraCmds_GatewayConfig, custom)
}

func BuildGatewayConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "config",
		Short:             `Gateway configuration management`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildGatewayConfigGetDefaultConfigurationCmd(),
		BuildGatewayConfigSetDefaultConfigurationCmd(),
		BuildGatewayConfigGetConfigurationCmd(),
		BuildGatewayConfigSetConfigurationCmd(),
		BuildGatewayConfigResetDefaultConfigurationCmd(),
		BuildGatewayConfigResetConfigurationCmd(),
		BuildGatewayConfigConfigurationHistoryCmd(),
	}, extraCmds_GatewayConfig...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

func BuildGatewayConfigGetDefaultConfigurationCmd() *cobra.Command {
	in := &driverutil.GetRequest{}
	cmd := &cobra.Command{
		Use:               "get-default",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetDefaultConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildGatewayConfigSetDefaultConfigurationCmd() *cobra.Command {
	in := &SetRequest{}
	cmd := &cobra.Command{
		Use:               "set-default",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				if curValue, err := client.GetDefaultConfiguration(cmd.Context(), &driverutil.GetRequest{}); err == nil {
					in.Spec = curValue
				}
				if edited, err := cliutil.EditInteractive(in.Spec); err != nil {
					return err
				} else {
					in.Spec = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.SetDefaultConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildGatewayConfigGetConfigurationCmd() *cobra.Command {
	in := &driverutil.GetRequest{}
	cmd := &cobra.Command{
		Use:               "get",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	return cmd
}

func BuildGatewayConfigSetConfigurationCmd() *cobra.Command {
	in := &SetRequest{}
	cmd := &cobra.Command{
		Use:               "set",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				if curValue, err := client.GetConfiguration(cmd.Context(), &driverutil.GetRequest{}); err == nil {
					in.Spec = curValue
				}
				if edited, err := cliutil.EditInteractive(in.Spec); err != nil {
					return err
				} else {
					in.Spec = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.SetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildGatewayConfigResetDefaultConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "reset-default",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.ResetDefaultConfiguration(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func BuildGatewayConfigResetConfigurationCmd() *cobra.Command {
	in := &ResetRequest{}
	cmd := &cobra.Command{
		Use:               "reset",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				if edited, err := cliutil.EditInteractive(in); err != nil {
					return err
				} else {
					in = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.ResetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildGatewayConfigConfigurationHistoryCmd() *cobra.Command {
	in := &driverutil.ConfigurationHistoryRequest{}
	cmd := &cobra.Command{
		Use:               "history",
		Short:             "",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := GatewayConfigContextInjector.ClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.ConfigurationHistory(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	cmd.RegisterFlagCompletionFunc("target", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ActiveConfiguration", "DefaultConfiguration"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func (in *SetRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SetRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Spec == nil {
		in.Spec = &GatewayConfigSpec{}
	}
	fs.AddFlagSet(in.Spec.FlagSet(append(prefix, "spec")...))
	return fs
}

func (in *SetRequest) RedactSecrets() {
	if in == nil {
		return
	}
	in.Spec.RedactSecrets()
}

func (in *SetRequest) UnredactSecrets(unredacted *SetRequest) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Spec.UnredactSecrets(unredacted.GetSpec()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "spec." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *GatewayConfigSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("GatewayConfigSpec", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Revision == nil {
		in.Revision = &v1.Revision{}
	}
	fs.AddFlagSet(in.Revision.FlagSet(append(prefix, "revision")...))
	if in.Server == nil {
		in.Server = &ServerSpec{}
	}
	fs.AddFlagSet(in.Server.FlagSet(append(prefix, "server")...))
	if in.Management == nil {
		in.Management = &ManagementServerSpec{}
	}
	fs.AddFlagSet(in.Management.FlagSet(append(prefix, "management")...))
	if in.Relay == nil {
		in.Relay = &RelayServerSpec{}
	}
	fs.AddFlagSet(in.Relay.FlagSet(append(prefix, "relay")...))
	if in.Health == nil {
		in.Health = &HealthServerSpec{}
	}
	fs.AddFlagSet(in.Health.FlagSet(append(prefix, "health")...))
	if in.Dashboard == nil {
		in.Dashboard = &DashboardServerSpec{}
	}
	fs.AddFlagSet(in.Dashboard.FlagSet(append(prefix, "dashboard")...))
	if in.Storage == nil {
		in.Storage = &StorageSpec{}
	}
	fs.AddFlagSet(in.Storage.FlagSet(append(prefix, "storage")...))
	if in.Certs == nil {
		in.Certs = &CertsSpec{}
	}
	fs.AddFlagSet(in.Certs.FlagSet(append(prefix, "certs")...))
	if in.Plugins == nil {
		in.Plugins = &PluginsSpec{}
	}
	fs.AddFlagSet(in.Plugins.FlagSet(append(prefix, "plugins")...))
	if in.Keyring == nil {
		in.Keyring = &KeyringSpec{}
	}
	fs.AddFlagSet(in.Keyring.FlagSet(append(prefix, "keyring")...))
	if in.AgentUpgrades == nil {
		in.AgentUpgrades = &AgentUpgradesSpec{}
	}
	fs.AddFlagSet(in.AgentUpgrades.FlagSet(append(prefix, "agent-upgrades")...))
	if in.RateLimiting == nil {
		in.RateLimiting = &RateLimitingSpec{}
	}
	fs.AddFlagSet(in.RateLimiting.FlagSet(append(prefix, "rate-limiting")...))
	if in.Auth == nil {
		in.Auth = &AuthSpec{}
	}
	fs.AddFlagSet(in.Auth.FlagSet(append(prefix, "auth")...))
	return fs
}

func (in *GatewayConfigSpec) RedactSecrets() {
	if in == nil {
		return
	}
	in.Storage.RedactSecrets()
	in.Certs.RedactSecrets()
	in.Auth.RedactSecrets()
}

func (in *GatewayConfigSpec) UnredactSecrets(unredacted *GatewayConfigSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Storage.UnredactSecrets(unredacted.GetStorage()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "storage." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Certs.UnredactSecrets(unredacted.GetCerts()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "certs." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Auth.UnredactSecrets(unredacted.GetAuth()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "auth." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *ServerSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ServerSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:8080"), &in.HttpListenAddress), strings.Join(append(prefix, "http-listen-address"), "."), "Address and port to serve the gateway's internal http server on.")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:9090"), &in.GrpcListenAddress), strings.Join(append(prefix, "grpc-listen-address"), "."), "Address and port to serve the gateway's external grpc server on.")
	return fs
}

func (in *ManagementServerSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ManagementServerSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:11080"), &in.HttpListenAddress), strings.Join(append(prefix, "http-listen-address"), "."), "Address and port to serve the management http server on.")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:11090"), &in.GrpcListenAddress), strings.Join(append(prefix, "grpc-listen-address"), "."), "Address and port to serve the management grpc server on.")
	return fs
}

func (in *RelayServerSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("RelayServerSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:11190"), &in.GrpcListenAddress), strings.Join(append(prefix, "grpc-listen-address"), "."), "Address and port to serve the relay grpc server on.")
	return fs
}

func (in *HealthServerSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("HealthServerSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:8086"), &in.HttpListenAddress), strings.Join(append(prefix, "http-listen-address"), "."), "Address and port to serve the gateway's internal health/metrics/profiling")
	return fs
}

func (in *DashboardServerSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("DashboardServerSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("0.0.0.0:12080"), &in.HttpListenAddress), strings.Join(append(prefix, "http-listen-address"), "."), "Address and port to serve the web dashboard on.")
	fs.Var(flagutil.StringPtrValue(nil, &in.Hostname), strings.Join(append(prefix, "hostname"), "."), "The hostname at which the dashboard is expected to be reachable. This is")
	fs.StringSliceVar(&in.TrustedProxies, strings.Join(append(prefix, "trusted-proxies"), "."), nil, "List of trusted proxies for the dashboard's http server.")
	return fs
}

func (in *StorageSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("StorageSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.EnumPtrValue(flagutil.Ptr(StorageType_Etcd), &in.Type), strings.Join(append(prefix, "type"), "."), "")
	if in.Etcd == nil {
		in.Etcd = &EtcdSpec{}
	}
	fs.AddFlagSet(in.Etcd.FlagSet(append(prefix, "etcd")...))
	if in.Jetstream == nil {
		in.Jetstream = &JetStreamSpec{}
	}
	fs.AddFlagSet(in.Jetstream.FlagSet(append(prefix, "jetstream")...))
	return fs
}

func (in *StorageSpec) RedactSecrets() {
	if in == nil {
		return
	}
	in.Etcd.RedactSecrets()
}

func (in *StorageSpec) UnredactSecrets(unredacted *StorageSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Etcd.UnredactSecrets(unredacted.GetEtcd()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "etcd." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *EtcdSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("EtcdSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringSliceVar(&in.Endpoints, strings.Join(append(prefix, "endpoints"), "."), nil, "")
	if in.Certs == nil {
		in.Certs = &MTLSSpec{}
	}
	fs.AddFlagSet(in.Certs.FlagSet(append(prefix, "certs")...))
	return fs
}

func (in *EtcdSpec) RedactSecrets() {
	if in == nil {
		return
	}
	in.Certs.RedactSecrets()
}

func (in *EtcdSpec) UnredactSecrets(unredacted *EtcdSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Certs.UnredactSecrets(unredacted.GetCerts()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "certs." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *MTLSSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("MTLSSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.ServerCA), strings.Join(append(prefix, "server-ca"), "."), "Path to the server CA certificate. Mutually exclusive with serverCAData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServerCAData), strings.Join(append(prefix, "server-ca-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded server CA certificate data. Mutually exclusive with serverCA.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientCA), strings.Join(append(prefix, "client-ca"), "."), "Path to the client CA certificate (not needed in all cases). Mutually exclusive with clientCAData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientCAData), strings.Join(append(prefix, "client-ca-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded client CA certificate data. Mutually exclusive with clientCA.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientCert), strings.Join(append(prefix, "client-cert"), "."), "Path to the certificate used for client-cert auth. Mutually exclusive with clientCertData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientCertData), strings.Join(append(prefix, "client-cert-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded client certificate data. Mutually exclusive with clientCert.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientKey), strings.Join(append(prefix, "client-key"), "."), "Path to the private key used for client-cert auth. Mutually exclusive with clientKeyData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientKeyData), strings.Join(append(prefix, "client-key-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded client key data. Mutually exclusive with clientKey.")
	return fs
}

func (in *MTLSSpec) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetServerCAData() != "" {
		in.ServerCAData = flagutil.Ptr("***")
	}
	if in.GetClientCAData() != "" {
		in.ClientCAData = flagutil.Ptr("***")
	}
	if in.GetClientCertData() != "" {
		in.ClientCertData = flagutil.Ptr("***")
	}
	if in.GetClientKeyData() != "" {
		in.ClientKeyData = flagutil.Ptr("***")
	}
}

func (in *MTLSSpec) UnredactSecrets(unredacted *MTLSSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetServerCAData() == "***" {
		if unredacted.GetServerCAData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "serverCAData"},
			})
		} else {
			*in.ServerCAData = *unredacted.ServerCAData
		}
	}
	if in.GetClientCAData() == "***" {
		if unredacted.GetClientCAData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "clientCAData"},
			})
		} else {
			*in.ClientCAData = *unredacted.ClientCAData
		}
	}
	if in.GetClientCertData() == "***" {
		if unredacted.GetClientCertData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "clientCertData"},
			})
		} else {
			*in.ClientCertData = *unredacted.ClientCertData
		}
	}
	if in.GetClientKeyData() == "***" {
		if unredacted.GetClientKeyData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "clientKeyData"},
			})
		} else {
			*in.ClientKeyData = *unredacted.ClientKeyData
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *JetStreamSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("JetStreamSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.Endpoint), strings.Join(append(prefix, "endpoint"), "."), "")
	fs.Var(flagutil.StringPtrValue(nil, &in.NkeySeedPath), strings.Join(append(prefix, "nkey-seed-path"), "."), "")
	return fs
}

func (in *CertsSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CertsSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.CaCert), strings.Join(append(prefix, "ca-cert"), "."), "Path to a PEM encoded CA certificate file. Mutually exclusive with caCertData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.CaCertData), strings.Join(append(prefix, "ca-cert-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded CA certificate data. Mutually exclusive with caCert.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServingCert), strings.Join(append(prefix, "serving-cert"), "."), "Path to a PEM encoded server certificate file. Mutually exclusive with servingCertData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServingCertData), strings.Join(append(prefix, "serving-cert-data"), "."), "\x1b[31m[secret]\x1b[0m PEM encoded server certificate data. Mutually exclusive with servingCert.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServingKey), strings.Join(append(prefix, "serving-key"), "."), "Path to a PEM encoded server key file. Mutually exclusive with servingKeyData.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServingKeyData), strings.Join(append(prefix, "serving-key-data"), "."), "\x1b[31m[secret]\x1b[0m String containing PEM encoded server key data. Mutually exclusive with servingKey.")
	return fs
}

func (in *CertsSpec) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetCaCertData() != "" {
		in.CaCertData = flagutil.Ptr("***")
	}
	if in.GetServingCertData() != "" {
		in.ServingCertData = flagutil.Ptr("***")
	}
	if in.GetServingKeyData() != "" {
		in.ServingKeyData = flagutil.Ptr("***")
	}
}

func (in *CertsSpec) UnredactSecrets(unredacted *CertsSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetCaCertData() == "***" {
		if unredacted.GetCaCertData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "caCertData"},
			})
		} else {
			*in.CaCertData = *unredacted.CaCertData
		}
	}
	if in.GetServingCertData() == "***" {
		if unredacted.GetServingCertData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "servingCertData"},
			})
		} else {
			*in.ServingCertData = *unredacted.ServingCertData
		}
	}
	if in.GetServingKeyData() == "***" {
		if unredacted.GetServingKeyData() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "servingKeyData"},
			})
		} else {
			*in.ServingKeyData = *unredacted.ServingKeyData
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *PluginsSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("PluginsSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("/var/lib/opni/plugins"), &in.Dir), strings.Join(append(prefix, "dir"), "."), "Directory to search for plugin binaries.")
	if in.Filters == nil {
		in.Filters = &PluginFilters{}
	}
	fs.AddFlagSet(in.Filters.FlagSet(append(prefix, "filters")...))
	if in.Cache == nil {
		in.Cache = &CacheSpec{}
	}
	fs.AddFlagSet(in.Cache.FlagSet(append(prefix, "cache")...))
	return fs
}

func (in *PluginFilters) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("PluginFilters", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringSliceVar(&in.Exclude, strings.Join(append(prefix, "exclude"), "."), nil, "")
	return fs
}

func (in *CacheSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CacheSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.EnumPtrValue(flagutil.Ptr(PatchEngine_Zstd), &in.PatchEngine), strings.Join(append(prefix, "patch-engine"), "."), "")
	fs.Var(flagutil.EnumPtrValue(flagutil.Ptr(CacheBackend_Filesystem), &in.Backend), strings.Join(append(prefix, "backend"), "."), "")
	if in.Filesystem == nil {
		in.Filesystem = &FilesystemCacheSpec{}
	}
	fs.AddFlagSet(in.Filesystem.FlagSet(append(prefix, "filesystem")...))
	return fs
}

func (in *FilesystemCacheSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("FilesystemCacheSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("/var/lib/opni/plugin-cache"), &in.Dir), strings.Join(append(prefix, "dir"), "."), "")
	return fs
}

func (in *KeyringSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("KeyringSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.StringSliceVar(&in.RuntimeKeyDirs, strings.Join(append(prefix, "runtime-key-dirs"), "."), nil, "Directories to search for files containing runtime keys.")
	return fs
}

func (in *AgentUpgradesSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("AgentUpgradesSpec", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Kubernetes == nil {
		in.Kubernetes = &KubernetesAgentUpgradeSpec{}
	}
	fs.AddFlagSet(in.Kubernetes.FlagSet(append(prefix, "kubernetes")...))
	return fs
}

func (in *KubernetesAgentUpgradeSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("KubernetesAgentUpgradeSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.EnumPtrValue(flagutil.Ptr(ImageResolverType_Kubernetes), &in.ImageResolver), strings.Join(append(prefix, "image-resolver"), "."), "")
	return fs
}

func (in *RateLimitingSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("RateLimitingSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.FloatPtrValue(flagutil.Ptr[float64](10.0), &in.Rate), strings.Join(append(prefix, "rate"), "."), "")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](50), &in.Burst), strings.Join(append(prefix, "burst"), "."), "")
	return fs
}

func (in *AuthSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("AuthSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.EnumValue(AuthSpec_Basic, &in.Kind), strings.Join(append(prefix, "kind"), "."), "")
	if in.Openid == nil {
		in.Openid = &OpenIDAuthSpec{}
	}
	fs.AddFlagSet(in.Openid.FlagSet(append(prefix, "openid")...))
	return fs
}

func (in *AuthSpec) RedactSecrets() {
	if in == nil {
		return
	}
	in.Openid.RedactSecrets()
}

func (in *AuthSpec) UnredactSecrets(unredacted *AuthSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Openid.UnredactSecrets(unredacted.GetOpenid()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "openid." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *OpenIDAuthSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("OpenIDAuthSpec", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.Issuer), strings.Join(append(prefix, "issuer"), "."), "The OP's Issuer identifier. This must exactly match the issuer URL")
	fs.Var(flagutil.StringPtrValue(nil, &in.CaCertData), strings.Join(append(prefix, "ca-cert-data"), "."), "Optional PEM-encoded CA certificate data for the issuer.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientId), strings.Join(append(prefix, "client-id"), "."), "")
	fs.Var(flagutil.StringPtrValue(nil, &in.ClientSecret), strings.Join(append(prefix, "client-secret"), "."), "\x1b[31m[secret]\x1b[0m ")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("sub"), &in.IdentifyingClaim), strings.Join(append(prefix, "identifying-claim"), "."), "IdentifyingClaim is the claim that will be used to identify the user")
	fs.StringSliceVar(&in.Scopes, strings.Join(append(prefix, "scopes"), "."), nil, "Scope specifies optional requested permissions.")
	return fs
}

func (in *OpenIDAuthSpec) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetClientSecret() != "" {
		in.ClientSecret = flagutil.Ptr("***")
	}
}

func (in *OpenIDAuthSpec) UnredactSecrets(unredacted *OpenIDAuthSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetClientSecret() == "***" {
		if unredacted.GetClientSecret() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "clientSecret"},
			})
		} else {
			*in.ClientSecret = *unredacted.ClientSecret
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *ResetRequest) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("ResetRequest", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Revision == nil {
		in.Revision = &v1.Revision{}
	}
	fs.AddFlagSet(in.Revision.FlagSet(prefix...))
	return fs
}