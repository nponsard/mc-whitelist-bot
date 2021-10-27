package messages

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	rcon "github.com/forewing/csgo-rcon"
	"github.com/nilsponsard/mc-whitelist-bot/internal/config"
	"github.com/nilsponsard/mc-whitelist-bot/pkg/verbosity"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func OnCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	conf := config.GetConfig()

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

	connError := false

	for _, rconAddr := range conf.Rcons {

		conn := rcon.New(rconAddr.Address, rconAddr.Password, time.Millisecond*500)
		output, err := conn.Execute("whitelist add " + username)

		if err != nil {
			connError = true
			verbosity.Error(err, output)
		}

		verbosity.Debug(output)
	}

	if connError {
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
	} else {
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}

}
