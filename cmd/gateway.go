package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	breakGlass   bool
	insecure     bool
	outputFormat string
	showDetails  bool
)

// NewGatewayCmd builds the 'gateway' command and its subcommands.
func NewGatewayCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in gateway command\n")
		out = io.Discard
	}

	//ValidArgs: []string{"set", "get", "update"},
	//Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	gatewayCmd := &cobra.Command{
		Use:   "gateway",
		Short: "Set, get or update the gateway host",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'set', 'get' or 'update' to manage the gateway host.")
		},
		Example: ` ibmcloud pag gateway set pag.us-south.appdomain.cloud --break-glass --insecure
  					ibmcloud pag gateway get --output json --details
  					ibmcloud pag gateway update`,
	}

	gatewayCmd.AddCommand(newGatewaySetCmd(out))
	gatewayCmd.AddCommand(newGatewayGetCmd(out))
	gatewayCmd.AddCommand(newGatewayUpdateCmd(out))

	return gatewayCmd
}

func newGatewaySetCmd(out io.Writer) *cobra.Command {

	gatewaySetCmd := &cobra.Command{
		Use:   "set [gateway_host_endpoint]",
		Short: "Set the gateway host",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			gatewayHostEndpoint := args[0]

			if breakGlass {
				fmt.Println("Break-glass mode enabled")
			}
			if insecure {
				fmt.Println("Insecure mode enabled")
			}

			// TODO : Implement Gateway set logic - Proxy service call here

			_, _ = fmt.Fprintf(out, "Setting gateway host with endpoint: %s\n", gatewayHostEndpoint)
			return nil
		},
	}

	gatewaySetCmd.Flags().BoolVar(&breakGlass, "break-glass", false, "Enable break-glass mode")
	gatewaySetCmd.Flags().BoolVar(&insecure, "insecure", false, "Allow insecure connections")

	return gatewaySetCmd
}

func newGatewayGetCmd(out io.Writer) *cobra.Command {

	gatewayGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Get the current gateway host",
		RunE: func(cmd *cobra.Command, args []string) error {
			if outputFormat != "" && outputFormat != "json" {
				return errors.New("invalid value for --output (only 'json' supported)")
			}

			if outputFormat == "json" {
				// print JSON
				fmt.Println("Print in json format")
			} else {
				// default plain output
				fmt.Println("Print in plain format")
			}
			return nil
		},
	}
	gatewayGetCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format (json)")
	gatewayGetCmd.Flags().BoolVar(&showDetails, "details", false, "Show detailed gateway information")

	return gatewayGetCmd
}

func newGatewayUpdateCmd(out io.Writer) *cobra.Command {

	updateGatewayCmd := &cobra.Command{
		Use:   "update",
		Short: "Update the gateway host",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Gateway update logic - Proxy service call here
			_, _ = fmt.Fprintf(out, "Updating gateway host\n")
			return nil
		},
	}
	return updateGatewayCmd
}
