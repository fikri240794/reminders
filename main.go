//go:generate go run github.com/google/wire/cmd/wire gen github.com/fikri240794/reminders/transports/restful_api/servers
//go:generate go run github.com/swaggo/swag/cmd/swag init --parseDependency --parseInternal
//go:generate go run github.com/swaggo/swag/cmd/swag fmt

package main

import (
	"github.com/fikri240794/reminders/cmd/database_migrations"
	"github.com/fikri240794/reminders/cmd/restful_api"
	"github.com/spf13/cobra"
)

var (
	cmd *cobra.Command
)

func init() {
	cmd = &cobra.Command{}

	cmd.AddCommand(restful_api.Cmd)
	cmd.AddCommand(database_migrations.Cmd)
}

func main() {
	cmd.Execute()
}
