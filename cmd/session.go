package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SessionCmd = &cobra.Command{}
var sessionDeleteCmd = &cobra.Command{}
var sessionDeleteUserCmd = &cobra.Command{}
var user string
var sessionID string

func init() {
	// Example `session` command
	SessionCmd = &cobra.Command{
		Use:   "session",
		Short: "Manage sessions",
	}

	// Subcommand: `session delete <session_id>`
	sessionDeleteCmd = &cobra.Command{
		Use:   "delete [session_id]",
		Short: "Delete a session",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sessionID = args[0]
			fmt.Printf("Deleting session with ID: %s\n", sessionID)
		},
	}

	// Subcommand: `session delete --user <user_id>`
	sessionDeleteUserCmd = &cobra.Command{
		Use:   "deleteuser",
		Short: "Delete sessions for a user",
		Run: func(cmd *cobra.Command, args []string) {
			if user == "" {
				fmt.Println("Error: --user flag is required")
				return
			}
			fmt.Printf("Deleting sessions for user: %s\n", user)
		},
	}
	sessionDeleteUserCmd.Flags().StringVar(&user, "user", "", "User ID")

	SessionCmd.AddCommand(sessionDeleteCmd)
	SessionCmd.AddCommand(sessionDeleteUserCmd)
	rootCmd.AddCommand(SessionCmd)
}
