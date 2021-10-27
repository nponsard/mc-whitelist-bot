package start

import (
	"os"
	"os/signal"
	"sync"

	"github.com/bwmarrin/discordgo"
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/mc-whitelist-bot/internal/config"
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

	// function to execute

	job.Action = func() {

		conf := config.GetConfig()

		// create a client

		discord, err := discordgo.New("Bot " + conf.Discord.Token)
		if err != nil {
			verbosity.Error(err)
			cli.Exit(1)
		}

		// Print a message when the bot is online

		discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {

			verbosity.Info("Connected as " + s.State.User.Username)
		})

		// register handler for messageCreate events

		discord.AddHandler(messages.OnCreate)

		// connect to discord servers

		err = discord.Open()
		if err != nil {
			verbosity.Error(err)
			cli.Exit(1)
		}

		// wait for an interrupt

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		verbosity.Debug("Gracefully shutdowning")

	}
}
