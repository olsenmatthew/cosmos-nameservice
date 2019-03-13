package cli

import (
	"fmt"

	"github.com/olsenmatthew/cosmos-nameservice/client/context"
	"github.com/olsenmatthew/cosmos-nameservice/codec"
	"github.com/cosmos/sdk-application-tutorial/x/nameservice"
	"github.com/spf13/cobra"
)

// GetCmdResolveNameQueries information about a name
func GetCmdResolveName(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		Use: "resolve [name]",
		Short: "resolve name",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", string(name))
				return nil
			}

			var out nameservice.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)

		},
	}
}

// GetCmdWhois queries information about a domain
func GetCmdWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		Use: "whois [name]",
		Short: "Query whois info of name",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", string(name))
				return nil
			}

			var out nameservice.Whois
			cddc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdNames queries a list of all name
func GetCmdNames(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		Use: "names",
		Short: "names",
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/names", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get any query names")
				return nil
			}

			var out nameservice.QueryResNames
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
