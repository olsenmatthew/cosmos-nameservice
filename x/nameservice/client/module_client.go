package client

import (
	"github.com/olsenmatthew/cosmos-nameservice/client"
	nameservicecmd "github.com/olsenmatthew/sdk-application-tutorial/x/nameservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey	string
	cdc			*amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient {storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group nameservice quirie under a subcommand
	namesvcQueryCmd := &cobra.Command {
		Use: "nameservice",
		Short: "Querying commands for the nameservice module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		nameservicecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return namesvcQueryCmd

}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	namesvcTxCmd := &cobra.Command {
		Use: "nameservice",
		Short: "Nameservice transaction subcommands",
	}

	namesvcTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.GetCmdBuyName(mc.cdc),
		nameservice.GetCmdSetName(mc.cdc),
	)...)

	return namesvcTxCmd

}
