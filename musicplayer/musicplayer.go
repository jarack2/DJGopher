package musicplayer

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// func init() {
// 	flag.StringVar(&token, "t", "airhorn.dca", "Bot Token")
// 	flag.Parse()
// }

var playing = false

var buffer = make([][]byte, 0)

func MusicPlayer(s *discordgo.Session, m *discordgo.MessageCreate, dir string) {

	if m.Content == "m!stop" {
		playMusic(s, m)
	} else {
		if m.Content == "m!gag" {
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, f := range files {
				err = loadSound(dir + f.Name())
				if err != nil {
					fmt.Println("Error loading sound: ", err)
					fmt.Println("Please copy $GOPATH/src/github.com/bwmarrin/examples/airhorn/airhorn.dca to this directory.")
					return
				}
			}
		} else{
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, f := range files {
				err = loadSound(dir + f.Name())
				err = loadSound("music/5SecSilence.dca")
				if err != nil {
					fmt.Println("Error loading sound: ", err)
					fmt.Println("Please copy $GOPATH/src/github.com/bwmarrin/examples/airhorn/airhorn.dca to this directory.")
					return
				}
			}
		}


		playMusic(s, m)
	}

	//// Create a new Discord session using the provided bot token.
	//dg, err := discordgo.New("Bot " + token)
	//if err != nil {
	//	fmt.Println("Error creating Discord session: ", err)
	//	return
	//}

	//// Register ready as a callback for the ready events.
	//dg.AddHandler(ready)
	//
	//// Register guildCreate as a callback for the guildCreate events.
	//dg.AddHandler(guildCreate)
	//
	//// We need information about guilds (which includes their channels),
	//// messages and voice states.
	//dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)
	//
	//// Open the websocket and begin listening.
	//err = dg.Open()
	//if err != nil {
	//	fmt.Println("Error opening Discord session: ", err)
	//}

	//// Wait here until CTRL-C or other term signal is received.
	//fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	//sc := make(chan os.Signal, 1)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	//<-sc
	//
	//// Cleanly close down the Discord session.
	//dg.Close()
}

//// This function will be called (due to AddHandler above) when the bot receives
//// the "ready" event from Discord.
//func ready(s *discordgo.Session, event *discordgo.Ready) {
//
//	// Set the playing status.
//	s.UpdateStatus(0, "!TayTay")
//}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func playMusic(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			return
		}

		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				err = playSound(s, g.ID, vs.ChannelID, m)
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}

				return
			}
		}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
//func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
//
//	if event.Guild.Unavailable {
//		return
//	}
//
//	for _, channel := range event.Guild.Channels {
//		if channel.ID == event.Guild.ID {
//			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
//			return
//		}
//	}
//}

// loadSound attempts to load an encoded sound file from disk.
func loadSound(path string) error {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string, m *discordgo.MessageCreate) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	if m.Content == "m!stop" {
		buffer = buffer[:0]
		vc.Disconnect()
		playing = false
	} else {
		// Sleep for a specified amount of time before playing the sound
		time.Sleep(250 * time.Millisecond)

		playing = true
		// Start speaking.
		vc.Speaking(true)

		// Send the buffer data.
		for _, buff := range buffer {
			vc.OpusSend <- buff
		}

		// Stop speaking
		vc.Speaking(false)

		// Sleep for a specificed amount of time before ending.
		time.Sleep(10000 * time.Millisecond)

		// Disconnect from the provided voice channel.
		vc.Disconnect()

		playing = false
	}


	return nil
}
