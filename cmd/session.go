package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewSessionCmd builds the 'session' command and its subcommands.
func NewSessionCmd(out io.Writer) *cobra.Command {
	if out == nil {
		out = io.Discard
	}

	sessionCmd := &cobra.Command{
		Use:   "session",
		Short: "Manage sessions",
	}

	sessionCmd.AddCommand(newSessionDeleteCmd(out))
	sessionCmd.AddCommand(newSessionDeleteUserCmd(out))
	sessionCmd.AddCommand(newSessionStatusCmd(out))

	return sessionCmd
}

func newSessionDeleteCmd(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [session_id]",
		Short: "Delete a session by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sessionID := args[0]
			// Ideally call into a service layer here.
			_, _ = fmt.Fprintf(out, "Deleting session with ID: %s\n", sessionID)
			return nil
		},
	}
}

func newSessionDeleteUserCmd(out io.Writer) *cobra.Command {
	var user string

	c := &cobra.Command{
		Use:   "deleteuser",
		Short: "Delete sessions for a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			if user == "" {
				return fmt.Errorf("--user flag is required")
			}
			_, _ = fmt.Fprintf(out, "Deleting sessions for user: %s\n", user)
			return nil
		},
	}

	c.Flags().StringVar(&user, "user", "", "User ID")
	return c
}

func newSessionStatusCmd(out io.Writer) *cobra.Command {
	var byUser bool

	c := &cobra.Command{
		Use:   "status [session_id]",
		Short: "Check session status or all sessions for a user",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if byUser {
				user, _ := cmd.Flags().GetString("user")
				if user == "" {
					return fmt.Errorf("--user flag is required when --by-user is set")
				}
				_, _ = fmt.Fprintf(out, "Status for all sessions of user %s: %s\n", user, "ACTIVE")
				return nil
			}

			if len(args) != 1 {
				return fmt.Errorf("expected session id when not using --by-user")
			}
			sessionID := args[0]
			_, _ = fmt.Fprintf(out, "Status of session %s: ACTIVE\n", sessionID)
			return nil
		},
	}

	c.Flags().BoolVar(&byUser, "by-user", false, "Show status for all sessions of a user")
	c.Flags().String("user", "", "User ID (required with --by-user)")
	return c
}
