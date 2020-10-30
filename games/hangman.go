package games

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	// "fmt"
)
var movesLeft = 11
const whitespace = "\t"
var testing = "765802303978340352" // discord testing channel
var chosen_word = "Keebler Elves"
var discord_friendly_guessed_word string
var guessed_word []byte
// var words [9]string = [9]string{ "dwarves", "buzzard", "Josh is an idiot", "buffoon", "xylophone", "espionage", "Taylor Swift is so hot", "I love neopets", "buffoon" }
func Hangman(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(testing, "Lets Play Hangman!")
	s.AddHandler(inputMessage)
	
	size := len(chosen_word)
	guessed_word := make([]byte, size)

	for i := 0; i < size; i++ { // creates the length of the string in underscores to print for hangman
		if chosen_word[i] == ' ' {
			guessed_word[i] = ' '
			discord_friendly_guessed_word += " "	
		} else {
			guessed_word[i] = '_'	
			discord_friendly_guessed_word += " \\_ "		
		}
	}
	
	// fmt.Println(string(guessed_word))
	s.ChannelMessageSend(testing, discord_friendly_guessed_word)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func inputMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if strings.Contains(discord_friendly_guessed_word, m.Content) {
		new_guess := discord_friendly_guessed_word
		for i := 0; i < len(chosen_word); i++ {
			if chosen_word[i] == ' ' {
				new_guess += " "	
			} else if string(chosen_word[i]) == m.Content {
				new_guess += m.Content
			} else {
				if guessed_word[i] == '_' {
					new_guess += " \\" + string(guessed_word[i]) + " "		
				} else {
					new_guess += " " + string(guessed_word[i]) + " "	
				}
			}
		discord_friendly_guessed_word = new_guess
		s.ChannelMessageSend(testing, discord_friendly_guessed_word)
	}
}
}