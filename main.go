package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/homeblest/pubg_stat_tracker/adding"
	"github.com/homeblest/pubg_stat_tracker/listing"
	"github.com/homeblest/pubg_stat_tracker/matches"
	"github.com/homeblest/pubg_stat_tracker/players"
	"github.com/homeblest/pubg_stat_tracker/requesting"
	"github.com/homeblest/pubg_stat_tracker/storage"
)

// Variables used for command line parameters
var (
	Token  string
	APIKey string
)

var addingSvc adding.Service
var requestSvc requesting.Service
var listingSvc listing.Service

func main() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&APIKey, "k", "", "PUBG API Key")
	flag.Parse()

	var playerStorage players.Repository
	var matchStorage matches.Repository

	// TODO: Migrate the repository to a DB model instead of in memory
	playerStorage = new(storage.MemoryPlayerStorage)

	addingSvc = adding.NewService(matchStorage, playerStorage)
	requestSvc = requesting.NewService(APIKey)
	listingSvc = listing.NewService(playerStorage)

	bot, err := discordgo.New("Bot " + Token)

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

	switch {
	case strings.HasPrefix(m.Content, "!hello"):
		s.ChannelMessageSend(m.ChannelID, "Hello friend!")
	case strings.HasPrefix(m.Content, "!stats"):
		player, err := listingSvc.GetPlayer("Homeblest")

		if err != nil {
			if err == players.ErrorPlayerNotFound {
				// Player wasn't found in storage, get him from the API and store him
				player, err := requestSvc.RequestPlayer("pc-eu", "Homeblest")
				if err != nil {
					fmt.Println(err)
					return
				}
				addingSvc.AddPlayer(*player)
			} else {
				fmt.Println(err)
				return
			}

		}

		playerString := fmt.Sprintf("I tried contacting the PUBG API, did it work? playerName: %s", player.Name)
		s.ChannelMessageSend(m.ChannelID, playerString)
	}

}
