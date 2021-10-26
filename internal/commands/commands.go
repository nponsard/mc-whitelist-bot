package commands

import (
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/mc-whitelist-bot/internal/commands/ping"
)

// configure subcommands
func SetupCommands(app *cli.Cli) {
	app.Command("ping", "ping", ping.Ping)
}
