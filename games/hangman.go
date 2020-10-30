package games

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"strconv"
	"unicode"
)

var dictionary [9]string = [9]string{ "dwarves", "buzzard", "Josh is an idiot", "buffoon", "xylophone", "espionage", "Taylor Swift is so hot", "I love neopets", "buffoon" }
var hangman_display [11]string = [11]string{ "  +---+\n|       |\n		|\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n|   |\n		|\n		|\n=========\n\n"}

var chosen_word string
var guessed_letters string
var guessed_word string

var testing = "765802303978340352" // discord testing channel

var movesLeft = 11
func Hangman(s *discordgo.Session, m *discordgo.MessageCreate) {
	chosen_word = dictionary[1]
	s.ChannelMessageSend(testing, "Lets Play Hangman!")
	s.ChannelMessageSend(testing, hangman_display[0])
	
	s.AddHandler(inputMessage)

	for i := 0; i < len(chosen_word); i++ { // creates the length of the string in underscores to print for hangman
		if chosen_word[i] == ' ' {
			guessed_word += " "	
		} else {
			guessed_word += "-"		
		}
	}
	
	s.ChannelMessageSend(testing, guessed_word)
}

// callback function for when the user guesses a letter
func inputMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if m.Author.ID == s.State.User.ID { // ignoring all messages created by the bot
		return
	}
	
	if (lost()) {
		return
	}
	
	user_input := m.Content
	
	if len(user_input) > 1 || !unicode.IsLetter(rune(user_input[0])) { // the user input should only be a single character and should be a letter [a-z]
		s.ChannelMessageSend(testing, "The user input is invalid. Please try again.\n")
		return
	}
	
	if strings.Contains(strings.ToLower(guessed_letters), strings.ToLower(user_input)) { // if the user has already guessed a letter
		s.ChannelMessageSend(testing, "You have already guessed that letter. Please try another one.\n")
		return
	} else {
			guessed_letters += strings.ToLower(user_input) + " "
			
			if strings.Contains(strings.ToLower(chosen_word), strings.ToLower(user_input)) { // check to see if the users input matches the word
				replaceWordWithSuccessfulGuess(user_input)
				s.ChannelMessageSend(testing, guessed_word)
				if won() {
					s.ChannelMessageSend(testing, "You guessed the word!\n" + "You can play again with the command: g!hangman restart\n" )
					return
				}
			} else { // chosen letter is not in the word
				movesLeft--
			
				if lost() { // the user has lost the game
					s.ChannelMessageSend(testing, "You failed to guess: " + chosen_word + ". Better luck next time.\n")
					s.ChannelMessageSend(testing, "You can play again with the command: g!hangman restart\n" )
				} else { // wrong input, but the user hasnt lost yet
					s.ChannelMessageSend(testing, user_input + " is not in the word.\n" + "You have " + strconv.Itoa(movesLeft) + " moves left.\n" + "Keep Guessing!\n" + guessed_word + "\n")
				}
			}
		
	}
}

func replace(str string, character rune, index int) string {
	result := []rune(str)
	result[index] = character
	return string(result)
}

func replaceWordWithSuccessfulGuess(user_input string) {
	new_guess := guessed_word
	for i := 0; i < len(chosen_word); i++ {
		if strings.ToLower(string(chosen_word[i])) == strings.ToLower(user_input) {
			char := rune(chosen_word[i])
			new_guess = replace(new_guess, char, i)
		}
	}
	guessed_word = new_guess
}

func lost() bool{
	if (movesLeft == 0) {
			return true
	}
	return false
}

func won() bool{
	if (chosen_word == guessed_word) {
			return true
	}
	return false
}

// TODO restart hangman game
// TODO count down moves when they get it wrong and display how close the man is to getting hanged
// TODO check if the user has already guessed a letter and do not count down moves
// TODO try to figure out how to put out a whole dictionary of dictionary
// TODO print out answer at end if they lose 
// TODO abstract handler to separate function
// TODO print letters they have used 