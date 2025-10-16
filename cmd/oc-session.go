package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewOCCmd builds the 'oc' command and its subcommands.
func NewOCCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in OC command\n")
		out = io.Discard
	}

	ocCmd := &cobra.Command{
		Use:   "oc",
		Short: "Openshift access to IBM Cloud Openshift Clusters",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'config' to download kubernetes config file.\nUse 'impersonate' to manage impersonation of clusters.")
		},
		Example: `  ibmcloud pag oc config <target-cluster> [--passcode <passcode>] [--apikey <apikey>] [--ticket-id <ticket-id>] [--endpoint <private|public|vpe>]
  					ibmcloud pag oc impersonate add <cluster name>
					ibmcloud pag oc impersonate delete <cluster name>
					ibmcloud pag oc impersonate list `,
	}
	ocCmd.AddCommand(dowloadOCConfigCmd(out))
	ocCmd.AddCommand(addImpersonateCmd(out))
	ocCmd.AddCommand(deleteImpersonateCmd(out))
	ocCmd.AddCommand(listImpersonateCmd(out))

	return ocCmd
}

func dowloadOCConfigCmd(out io.Writer) *cobra.Command {

	var (
		isTicketIDProvided string
		isEndpointProvided string
		isPasscodeProvided string
		isAPIKeyProvided   string
	)
	downloadOCConfigCmd := &cobra.Command{
		Use:   "config target_cluster",
		Short: "Configures the user's kubernetes client to access a given kubernetes cluster.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(out, "Configuring kubernetes client for cluster:", args[0], "\n.")

			if isTicketIDProvided != "" {
				fmt.Fprintln(out, "Ticket-ID provided:", isTicketIDProvided, "\n.")
			}
			if isEndpointProvided != "" {
				fmt.Fprintln(out, "Endpoint provided:", isEndpointProvided, "\n.")
			}
			if isPasscodeProvided != "" {
				fmt.Fprintln(out, "Passcode provided:", isPasscodeProvided, "\n.")
			}
			if isAPIKeyProvided != "" {
				fmt.Fprintln(out, "APIKey provided:", isAPIKeyProvided, "\n.")
			}
			// We need custome logic to see the combination of flags...
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	downloadOCConfigCmd.Flags().StringVar(&isTicketIDProvided, "ticket-id", "", "Ticket ID for the session")
	downloadOCConfigCmd.Flags().StringVar(&isEndpointProvided, "endpoint", "", "private, public or vpe")
	downloadOCConfigCmd.Flags().StringVar(&isPasscodeProvided, "passcode", "", "Passcode for the cluster")
	downloadOCConfigCmd.Flags().StringVar(&isAPIKeyProvided, "apikey", "", "APIKey for the cluster")

	return downloadOCConfigCmd
}
