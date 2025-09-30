package main

import (
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/yathendra/cobrapagcliplugin/cmd"
)

type pagPlugin struct{}

func (p *pagPlugin) Run(context plugin.PluginContext, args []string) {
	if len(args) == 0 {
		terminal.NewStdUI().Warn("No command provided")
		return
	}

	pagCmd := cmd.NewPagCmd(os.Stdout)
	pagCmd.SetArgs(args)
	if err := pagCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func (p *pagPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    "pag",
		Version: plugin.VersionType{Major: 0, Minor: 1, Build: 0},
		Namespaces: []plugin.Namespace{
			{
				Name:        "pag",
				Description: "PAG with cobra plugin POC.",
			},
		},
		Commands: []plugin.Command{
			{
				Namespace:   "pag",
				Name:        "session",
				Description: "Session Management commands",
				Usage:       "ibmcloud pag session",
			},
			{
				Namespace:   "pag",
				Name:        "gateway",
				Description: "Gateway Management commands",
				Usage:       "ibmcloud pag gateway",
			},
			{
				Namespace:   "pag",
				Name:        "gateway set",
				Description: "Set the gateway host URL",
				Usage:       "ibmcloud pag gateway set <gateway url> [--break-glass] [--insecure]",
			},
			{
				Namespace:   "pag",
				Name:        "gateway get",
				Description: "Get the gateway host URL",
				Usage:       "ibmcloud pag gateway get [--output <json>] [--details]",
			},
			{
				Namespace:   "pag",
				Name:        "gateway update",
				Description: "Update PAG gateway to the latest proxy image",
				Usage:       "ibmcloud pag gateway update",
			},
		},
	}
}

func main() {
	plugin.Start(new(pagPlugin))
}
