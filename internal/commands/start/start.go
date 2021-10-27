package start

import (
	"os"
	"os/signal"
	"sync"

	"github.com/bwmarrin/discordgo"
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/mc-whitelist-bot/internal/messages"
	"github.com/nilsponsard/mc-whitelist-bot/pkg/verbosity"
)

var (
	mapLock      sync.Mutex
	rconPassword *string
	queueMap     map[string]chan int
)

// setup ping command
func Start(job *cli.Cmd) {

	queueMap = make(map[string]chan int)

	token := job.StringArg("TOKEN", "", "Discord token")
	rconPassword = job.StringArg("RCON_PASSWORD", "", "RCON password")

	// function to execute

	job.Action = func() {

		discord, err := discordgo.New("Bot " + *token)
		if err != nil {
			verbosity.Error(err)
			cli.Exit(1)
		}

		discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
			verbosity.Info("Bot is up!")
		})
		discord.AddHandler(messages.OnCreate)

		err = discord.Open()
		if err != nil {
			verbosity.Error(err)
			cli.Exit(1)
		}

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		verbosity.Debug("Gracefully shutdowning")

	}
}
