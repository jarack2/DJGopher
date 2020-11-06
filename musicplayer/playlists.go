package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

// Folder is folder
var (
	Folder = flag.String("f", "/mnt/c/Users/Jared/Downloads/MUSIC", "Folder of files to play.")
	err    error
)

// MusicPlayer plays music
func MusicPlayer(s *discordgo.Session, guildID, channelID string) {
	// Connect to voice channel.
	// NOTE: Setting mute to false, deaf to true.
	dgv, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start loop and attempt to play all files in the given folder
	fmt.Println("Reading Folder: ", *Folder)
	files, _ := ioutil.ReadDir(*Folder)
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())
		s.UpdateStatus(0, f.Name())

		dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", *Folder, f.Name()), make(chan bool))
	}

	// Close connections
	dgv.Close()

	return
}
