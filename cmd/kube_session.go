package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewKSCmd builds the 'ks' command and its subcommands.
func NewKSCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in KS command\n")
		out = io.Discard
	}

	ksCmd := &cobra.Command{
		Use:   "ks",
		Short: "Kubernetes access to IBM Cloud Kubernetes Clusters",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'config' to download kubernetes config file.\nUse 'impersonate' to manage impersonation of clusters.")
		},
		Example: `  ibmcloud pag ks config <target-cluster> [--ticket-id <ticket-id>] [--endpoint <private|public|vpe>]
  					ibmcloud pag ks impersonate add <cluster name>
					ibmcloud pag ks impersonate delete <cluster name>
					ibmcloud pag ks impersonate list `,
	}

	ksCmd.AddCommand(dowloadKSConfigCmd(out))
	ksCmd.AddCommand(addImpersonateCmd(out))
	ksCmd.AddCommand(deleteImpersonateCmd(out))
	ksCmd.AddCommand(listImpersonateCmd(out))

	return ksCmd
}

func dowloadKSConfigCmd(out io.Writer) *cobra.Command {

	var (
		isTicketIDProvided string
		isEndpointProvided string
	)
	downloadKSConfigCmd := &cobra.Command{
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
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	downloadKSConfigCmd.Flags().StringVar(&isTicketIDProvided, "ticket-id", "", "Ticket ID for the session")
	downloadKSConfigCmd.Flags().StringVar(&isEndpointProvided, "endpoint", "", "private, public or vpe")

	return downloadKSConfigCmd
}

func addImpersonateCmd(out io.Writer) *cobra.Command {

	addImpersonateCmd := &cobra.Command{
		Use:   "impersonate add cluster name",
		Short: "Enables PAG to communicate directly with the Kubernetes cluster",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Adding impersonate for cluster:", args[0], "\n.")
			return nil
		},
	}

	return addImpersonateCmd
}

func deleteImpersonateCmd(out io.Writer) *cobra.Command {

	deleteImpersonateCmd := &cobra.Command{
		Use:   "impersonate delete cluster name",
		Short: "Cleanup IKS impersonation resources from a cluster",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Deleting impersonate for cluster:", args[0], "\n.")
			return nil
		},
	}

	return deleteImpersonateCmd
}

func listImpersonateCmd(out io.Writer) *cobra.Command {

	listImpersonateCmd := &cobra.Command{
		Use:   "impersonate list",
		Short: "List impersonated clusters",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Listing impersonated clusters\n.")
			return nil
		},
	}

	return listImpersonateCmd
}
