package games

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	// "fmt"
)
var movesLeft = 11
const whitespace = "\t"
var testing = "765802303978340352" // discord testing channel
var chosen_word string
var discord_friendly_guessed_word string
var guessed_word []byte
var words [9]string = [9]string{ "dwarves", "buzzard", "Josh is an idiot", "buffoon", "xylophone", "espionage", "Taylor Swift is so hot", "I love neopets", "buffoon" }
func Hangman(s *discordgo.Session, m *discordgo.MessageCreate) {
	start_string := "  +---+ \n  |   |\n      |\n      |\n      |\n      |\n========="
	chosen_word = words[6]
	s.ChannelMessageSend(testing, "Lets Play Hangman!")
	s.ChannelMessageSend(testing, start_string)
	
	s.AddHandler(inputMessage)
	
	size := len(chosen_word)
	guessed_word := make([]byte, size)

	for i := 0; i < size; i++ { // creates the length of the string in underscores to print for hangman
		if chosen_word[i] == ' ' {
			guessed_word[i] = ' '
			discord_friendly_guessed_word += " "	
		} else {
			guessed_word[i] = '-'	
			discord_friendly_guessed_word += "-"		
		}
	}
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
	if strings.Contains(strings.ToLower(chosen_word), strings.ToLower(m.Content)) {
		new_guess := discord_friendly_guessed_word
		for i := 0; i < len(chosen_word); i++ {
			if strings.ToLower(string(chosen_word[i])) == strings.ToLower(m.Content) {
				char := rune(chosen_word[i])
				new_guess = replace(new_guess, char, i)
			}
		}
		discord_friendly_guessed_word = new_guess
		s.ChannelMessageSend(testing, discord_friendly_guessed_word)
	}
}

func replace(str string, character rune, index int) string {
	result := []rune(str)
	result[index] = character
	return string(result)
}

func convertByteArray(in []byte) string {
	result := ""
	for i := 0; i < len(in); i++ {
		if in[i] == ' ' {
			
		}
	}
	return result
}

// TODO restart hangman game
// TODO count down moves when they get it wrong and display how close the man is to getting hanged
// TODO check if the user has already guessed a letter and do not count down moves
// TODO try to figure out how to put out a whole dictionary of words
// TODO print out answer at end if they lose 
// TODO abstract handler to separate function