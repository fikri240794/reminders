package restful_api

import (
	"github.com/fikri240794/reminders/configs"
	"github.com/fikri240794/reminders/transports/restful_api/servers"
	"github.com/spf13/cobra"
)

var (
	restfulapi_server servers.IRESTfulAPIServer
	cfgPath           string
	Cmd               *cobra.Command
)

func preRun(cmd *cobra.Command, args []string) {
	configs.SetConfigPath(cfgPath)
	restfulapi_server = servers.NewRESTfulAPIServer()
}

func run(cmd *cobra.Command, args []string) {
	restfulapi_server.Serve()
}

func init() {
	Cmd = &cobra.Command{
		Use:    "restful-api-server",
		PreRun: preRun,
		Run:    run,
	}
	Cmd.PersistentFlags().StringVar(
		&cfgPath,
		"config-path",
		".env",
		"config file path",
	)
}
