package main

import (
	"./games"
	"./musicplayer"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const Token string = "NzcwMDAyMzExODc4OTM0NTI4.X5XOiQ.Z9F3_0y55l_VScYv7qx_zbV38rg"

var game_running = false
var trivia_game_running = false
var BotID string

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the runProgram func as a callback for MessageCreate events.
	dg.AddHandler(runProgram)
	dg.AddHandler(guildCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func runProgram(s *discordgo.Session, m *discordgo.MessageCreate) {
	testing := "765802303978340352"
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "g!ping" {
		s.ChannelMessageSend(testing, "Pong!")
		return
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "g!pong" {
		s.ChannelMessageSend(testing, "Ping!")
		return
	}

	// If the message is "pog" reply with ":gitpog:"
	if m.Content == "pog" {
		s.ChannelMessageSend(testing, "<:gitpog:770159988915044352>")
		return
	}

	if m.Content == "g!hangman restart" {
		games.Restart(s, m)
		game_running = false
		return
	}

	if m.Content == "g!hangman" || game_running == true {
		if !game_running {
			games.Hangman(s, m, game_running)
			game_running = true
			return
		}
		games.Hangman(s, m, game_running)
		return
	}

	if m.Content == "g!trivia" || trivia_game_running == true {
		if trivia_game_running {
			games.Trivia(s, m, trivia_game_running)
			trivia_game_running = true
			return
		}
		trivia_game_running = games.Trivia(s, m, trivia_game_running)
		return
	}

	if m.Content == "sup" {
		s.ChannelMessageSend(testing, "sup")
		return
	}

	if m.Content == "m!play" {
		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			fmt.Println(err)
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			fmt.Println(err)
			return
		}

		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				music.MusicPlayer(s, g.ID, vs.ChannelID)
				return
			}
		}
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}
