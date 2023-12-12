//go:build !minimal

package commands

import (
	"sync"

	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/clients"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/tracing"
	"github.com/rancher/opni/pkg/util/flagutil"
	"github.com/spf13/cobra"
)

var mgmtClient managementv1.ManagementClient
var managementListenAddress string
var initManagementOnce sync.Once
var lg = logger.New()

func ConfigureManagementCommand(cmd *cobra.Command) {
	if cmd.PersistentPreRunE == nil {
		cmd.PersistentPreRunE = managementPreRunE
	} else {
		oldPreRunE := cmd.PersistentPreRunE
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			if err := managementPreRunE(cmd, args); err != nil {
				return err
			}
			return oldPreRunE(cmd, args)
		}
	}
}

func managementPreRunE(cmd *cobra.Command, _ []string) (err error) {
	initManagementOnce.Do(func() {
		tracing.Configure("cli")
		address := cmd.Flag("address").Value.String()
		if address == "" {
			var mgmtConfig configv1.ManagementServerSpec
			flagutil.LoadDefaults(&mgmtConfig)
			address = mgmtConfig.GetGrpcListenAddress()
		}
		managementListenAddress = address
		mgmtClient, err = clients.NewManagementClient(cmd.Context(), clients.WithAddress(address))
	})
	return err
}
