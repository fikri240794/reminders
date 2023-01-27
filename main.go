package main

import (
	"github.com/fikri240794/reminders/cmd/database_migrations"
	"github.com/fikri240794/reminders/cmd/restful_api_servers"
	"github.com/spf13/cobra"
)

var (
	cmd *cobra.Command
)

func init() {
	cmd = &cobra.Command{}

	cmd.AddCommand(restful_api_servers.Cmd)
	cmd.AddCommand(database_migrations.Cmd)
}

func main() {
	cmd.Execute()
}
