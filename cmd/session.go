package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewSessionCmd builds the 'session' command and its subcommands.
func NewSessionCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in session command\n")
		out = io.Discard
	}

	sessionCmd := &cobra.Command{
		Use:   "session",
		Short: "Session Management commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'session list' to list active sessions.")
		},
		Example: `ibmcloud pag session list`,
	}

	sessionCmd.AddCommand(listSessionCmd(out))
	return sessionCmd
}

func NewSessionsCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in sessions command\n")
		out = io.Discard
	}

	sessionsCmd := &cobra.Command{
		Use:   "sessions",
		Short: "Session Management commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Listing all active sessions\n.")
			return nil
		},
		Example: `ibmcloud pag sessions`,
	}

	sessionsCmd.AddCommand(listSessionCmd(out))
	return sessionsCmd
}

func listSessionCmd(out io.Writer) *cobra.Command {

	listImpersonateCmd := &cobra.Command{
		Use:   "list",
		Short: "Display list of sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO : Implement Proxy service call here
			fmt.Fprintln(out, "Listing all active sessions\n.")
			return nil
		},
	}

	return listImpersonateCmd
}
