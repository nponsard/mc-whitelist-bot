package start

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	rcon "github.com/forewing/csgo-rcon"
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/mc-whitelist-bot/pkg/verbosity"
)

var (
	mapLock  sync.Mutex
	queueMap map[string]chan int
)

// setup ping command
func Start(job *cli.Cmd) {

	queueMap = make(map[string]chan int)

	token := job.StringArg("TOKEN", "", "Discord token")

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
		discord.AddHandler(messageCreate)

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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	verbosity.Debug(m.ChannelID)

	msgChannel, err := s.State.Channel(m.ChannelID)

	if err != nil {
		verbosity.Error(err)
		return
	}
	if msgChannel.Name != "whitelist" {
		return
	}

	username := strings.Trim(m.Content, " \t\n")

	conn := rcon.New("127.0.0.1:25575", "jesuissautax", time.Millisecond*50)

	output, err := conn.Execute("whitelist add " + username)
	if err != nil {
		verbosity.Error(err, output)
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
	} else {
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
	verbosity.Debug(output)
}
