package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pag",
	Short: "IBM Cloud CLI plugin example",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
