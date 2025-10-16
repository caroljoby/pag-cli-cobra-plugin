package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewBreakGlassCmd builds the 'break-glass' command and its subcommands.
func NewBreakGlassCmd(out io.Writer) *cobra.Command {
	if out == nil {
		fmt.Print("nil out in break-glass command\n")
		out = io.Discard
	}

	breakGlassCmd := &cobra.Command{
		Use:   "break-glass",
		Short: "Break Glass Actions",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "Use 'key' to manage break-glass keys.\n . Use 'keys' to list all break-glass keys.")
		},
		Example: `  ibmcloud pag break-glass key create
  					ibmcloud pag break-glass key delete <user IAM id>
					ibmcloud pag break-glass key get --help
					ibmcloud pag break-glass key list --help
					ibmcloud pag break-glass keys `,
	}

	breakGlassCmd.AddCommand(breakGlassKeyCreateCmd(out))
	breakGlassCmd.AddCommand(breakGlassKeyDeleteCmd(out))
	breakGlassCmd.AddCommand(breakGlassKeyGetCmd(out))
	breakGlassCmd.AddCommand(breakGlassKeysListCmd(out))

	return breakGlassCmd
}

func breakGlassKeyCreateCmd(out io.Writer) *cobra.Command {

	breakGlassKeyCreateCmd := &cobra.Command{
		Use:   "key create",
		Short: "Create break-glass keys for the current user and register.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get user IAM ID from ACESS_TOKEN
			fmt.Fprintln(out, "Creating break-glass key for user:", args[0], "\n.")
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	return breakGlassKeyCreateCmd
}

func breakGlassKeyDeleteCmd(out io.Writer) *cobra.Command {

	breakGlassKeyDeleteCmd := &cobra.Command{
		Use:   "key delete [user IAM_ID]",
		Short: "Delete the break-glass key for the current user. If you have manager access, you can delete other user keys using the --user <IAMId> option.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get user IAM ID from ACESS_TOKEN
			fmt.Fprintln(out, "Deleting break-glass key for user:")
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	return breakGlassKeyDeleteCmd
}

func breakGlassKeyGetCmd(out io.Writer) *cobra.Command {

	breakGlassKeyGetCmd := &cobra.Command{
		Use:   "key get",
		Short: "Get the break-glass key info for the current use.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get user IAM ID from ACESS_TOKEN
			fmt.Fprintln(out, "Getting break-glass key for user:", args[0], "\n.")
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	return breakGlassKeyGetCmd
}

func breakGlassKeysListCmd(out io.Writer) *cobra.Command {

	breakGlassKeysListCmd := &cobra.Command{
		Use:   "keys list",
		Short: "List all break-glass keys.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get user IAM ID from ACESS_TOKEN
			fmt.Fprintln(out, "Listing all break-glass keys\n.")
			// TODO : Implement Proxy service call here
			return nil
		},
	}

	return breakGlassKeysListCmd
}
