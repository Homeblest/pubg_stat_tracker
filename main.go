package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/homeblest/pubg_stat_tracker/adding"
	"github.com/homeblest/pubg_stat_tracker/matches"
	"github.com/homeblest/pubg_stat_tracker/players"
	"github.com/homeblest/pubg_stat_tracker/storage"
)

func main() {
	var playerStorage players.Repository
	var matchStorage matches.Repository

	// TODO: Migrate the repository to a DB model instead of in memory
	playerStorage = new(storage.MemoryPlayerStorage)

	addingSvc := adding.New(matchStorage, playerStorage)

	player := players.Player{
		ID:      "1234",
		Name:    "hjalti",
		Region:  "eu",
		Matches: nil,
	}

	addingSvc.AddPlayer(player)

	// TODO: Move the bot token to be a flag argument when booting
	bot, err := discordgo.New("Bot NDU0NjQ2NzU2MDU2MTA0OTYy.DfweXQ.NxksM0kHOwr4sQJuoys6E_2PiRQ")

	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("PUBG_Stat_Tracker is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!hello") {
		s.ChannelMessageSend(m.ChannelID, "Hello friend!")
	}
}
