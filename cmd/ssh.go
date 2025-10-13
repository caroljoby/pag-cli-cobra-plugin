package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	isTicketIDEnabled bool
)

// NewSSHCmd builds the 'ssh' command and its subcommands.
func NewSSHCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in SSH command\n")
		out = io.Discard
	}

	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "SSH access to Linux Virtual Server Instances in an IBM Cloud VPC",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'connect' to access Virtual server instance.")
		},
		Example: `  ibmcloud pag ssh ca
  					ibmcloud pag ssh cert get [--ticket-id <ticket-id>]
  					ibmcloud pag ssh connect <virtual server address> [--ticket-id <ticket-id>]`,
	}

	sshCmd.AddCommand(getSSHCaCmd(out))
	sshCmd.AddCommand(getSSHCertDownloadCmd(out))
	sshCmd.AddCommand(getSSHConnectCmd(out))

	return sshCmd
}

func getSSHCaCmd(out io.Writer) *cobra.Command {

	sshCaCmd := &cobra.Command{
		Use:   "ca",
		Short: "Retrieve the CA Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Retrieving CA Keys\n.")
			return nil
		},
	}
	return sshCaCmd
}

func getSSHCertDownloadCmd(out io.Writer) *cobra.Command {

	sshCertDownloadCmd := &cobra.Command{
		Use:   "cert get",
		Short: "Download certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Downloading certificate\n.")

			if isTicketIDEnabled {
				fmt.Println("Ticket-ID mode enabled")
			}
			return nil
		},
	}

	sshCertDownloadCmd.Flags().BoolVar(&isTicketIDEnabled, "ticket-id", false, "Ticket ID for the session")

	return sshCertDownloadCmd
}

func getSSHConnectCmd(out io.Writer) *cobra.Command {

	sshConnectCmd := &cobra.Command{
		Use:   "connect [virtual_server_instance_address]",
		Short: "Connect to new ssh session for a virtual server instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintf(out, "Connecting to virtual server instance with address: %s.", args[0])
			if isTicketIDEnabled {
				fmt.Println("Ticket-ID mode enabled")
			}
			// TODO : Implement Gateway set logic - Proxy service call here
			return nil
		},
	}

	sshConnectCmd.Flags().BoolVar(&isTicketIDEnabled, "ticket-id", false, "Ticket ID for the session")

	return sshConnectCmd
}
