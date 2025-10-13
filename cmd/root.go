package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// NewPagCmd returns a configured root *cobra.Command.
// Provide an io.Writer for output (useful for tests); pass os.Stdout in real use.
func NewPagCmd(out io.Writer) *cobra.Command {
	if out == nil {
		out = os.Stdout
	}

	pagCmd := &cobra.Command{
		Use:     "pag",
		Short:   "IBM Cloud PAG CLI plugin",
		Long:    `PAG CLI Plugin for IBM CLOUD.`,
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Welcome to the PAG CLI ! Use -h to see available commands.")
		},
	}

	// register subcommands
	pagCmd.AddCommand(NewSessionCmd(out))
	pagCmd.AddCommand(NewGatewayCmd(out))
	pagCmd.AddCommand(NewSSHCmd(out))

	return pagCmd
}
