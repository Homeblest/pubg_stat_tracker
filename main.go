package main

import (
	"bytes"
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
	"github.com/homeblest/pubg_stat_tracker/regions"
	"github.com/homeblest/pubg_stat_tracker/requesting"
	"github.com/homeblest/pubg_stat_tracker/storage"
)

var (
	token  string
	apiKey string
)

var addingSvc adding.Service
var requestSvc requesting.Service
var listingSvc listing.Service

func main() {
	apiKey = os.Getenv("PUBG_API_KEY")
	token = os.Getenv("DISCORD_BOT_TOKEN")

	var playerStorage players.Repository
	var matchStorage matches.Repository

	// TODO: Migrate the repository to a DB model instead of in memory
	playerStorage = new(storage.MemoryPlayerStorage)

	addingSvc = adding.NewService(matchStorage, playerStorage)
	requestSvc = requesting.NewService(apiKey)
	listingSvc = listing.NewService(playerStorage)

	bot, err := discordgo.New("Bot " + token)

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
	// Ignore all messages created by the bot itself, and if the message doesn't start with ! (bang)
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, "!") {
		return
	}

	inputString := strings.Split(m.Content, " ")
	cmd := inputString[0]

	switch cmd {
	case "!status":
		status, err := requestSvc.RequestStatus()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Apologies, the PUBG API seems to be down at the moment")
			return
		}
		statusString := fmt.Sprintf("Type: %s - ID: %s - Release Date: %s - Version: %s",
			status.Data.Type, status.Data.ID, status.Data.Attributes.ReleasedAt, status.Data.Attributes.Version,
		)
		s.ChannelMessageSend(m.ChannelID, statusString)
	case "!stats":
		if len(inputString) < 3 {
			s.ChannelMessageSend(m.ChannelID, "Invalid command, use: !stats <playerName> <region> <gameMode>")
			return
		}
		playerName := inputString[1]
		shardID := regions.GetShardIDFromRegion(inputString[2])
		// gameMode := inputString[3]

		player, err := requestSvc.RequestPlayer(playerName, shardID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		allSeasons, err := requestSvc.RequestSeasons(shardID)
		if err != nil {
			fmt.Println(err.Error())
			s.ChannelMessageSend(m.ChannelID, "Failed requesting seasons")
			return
		}
		var currentSeasonID string

		for _, season := range allSeasons.Seasons {
			if season.Attributes.IsCurrentSeason {
				currentSeasonID = season.ID
			}
		}

		playerSeasonStats, err := requestSvc.RequestSeasonStatistics(player.ID, currentSeasonID, shardID)
		statisticsString := fmt.Sprintf("You have killed %d players", playerSeasonStats.Data.Attributes.GameModeStats.SquadFPP.Kills)

		s.ChannelMessageSend(m.ChannelID, statisticsString)
	case "!seasons":
		seasonData, err := requestSvc.RequestSeasons("pc-eu")
		if err != nil {
			fmt.Println(err)
			return
		}
		var buffer bytes.Buffer
		for _, season := range seasonData.Seasons {
			buffer.WriteString(fmt.Sprintf("%s\n", season.ID))
		}
		s.ChannelMessageSend(m.ChannelID, buffer.String())

	}

}
