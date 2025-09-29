package main

import (
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

	switch args[0] {
	case "session":
		cmd.SessionCmd.Execute()
	default:
		terminal.NewStdUI().Warn("Unknown command: " + args[0])
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
			// {
			// 	Namespace:   "pag",
			// 	Name:        "goodbye",
			// 	Description: "Print goodbye message",
			// 	Usage:       "ibmcloud pag goodbye",
			// },
		},
	}
}

func main() {
	plugin.Start(new(pagPlugin))
}
