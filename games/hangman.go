package games

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"strconv"
	//"unicode"
)

var dictionary [9]string = [9]string{ "dwarves", "buzzard", "buffoon", "xylophone", "espionage", "Taylor Swift is so hot", "I love neopets", "buffoon" } // a makeshift dictionary
var hangman_display [7]string = [7]string{ "  +---+\n|       |\n		|\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n|   |\n		|\n		|\n=========\n\n",  "  +---+\n|       |\n		|\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n		|\n		|\n		|\n=========\n\n", "  +---+\n|   |\nO   |\n|   |\n		|\n		|\n=========\n\n"}

var chosen_word string // the word picked from the dictionary
var guessed_letters string // the letters that the user has guessed
var guessed_word string // the word that the user is currently testing

var testing = "765802303978340352" // discord testing channel

var movesLeft = 7 // how many wrong moves the user has left
var display = 0 // where the display is in the array
func Hangman(s *discordgo.Session, m *discordgo.MessageCreate, game_running bool) {
	if !game_running {
		chosen_word = dictionary[1]
		s.ChannelMessageSend(testing, "Lets Play Hangman!")
		s.ChannelMessageSend(testing, hangman_display[0])
	
		for i := 0; i < len(chosen_word); i++ { // creates the length of the string in underscores to print for hangman
			if chosen_word[i] == ' ' {
				guessed_word += " "	
			} else {
				guessed_word += "-"		
			}
		}
		
		s.ChannelMessageSend(testing, guessed_word)
	} else {
		inputMessage(s, m)
	}
	return
}

// callback function for when the user guesses a letter
func inputMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if m.Author.ID == s.State.User.ID { // ignoring all messages created by the bot
		return
	}
	
	if (lost() || won()) {
		return
	}
	
	user_input := m.Content
	
	//if len(user_input) > 1 || !unicode.IsLetter(rune(user_input[0])) { // the user input should only be a single character and should be a letter [a-z]
	//	s.ChannelMessageSend(testing, "The user input is invalid. Please try again.\n")
	//	return
	//}
	
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
				return
			} else { // chosen letter is not in the word
				movesLeft--
			
				if lost() { // the user has lost the game
					s.ChannelMessageSend(testing, "You failed to guess: \"" + chosen_word + "\". Better luck next time.\n")
					s.ChannelMessageSend(testing, "You can play again with the command: g!hangman restart\n" )
					return
				} else { // wrong input, but the user hasnt lost yet
					display++
					s.ChannelMessageSend(testing, user_input + " is not in the word.\n" + "You have " + strconv.Itoa(movesLeft) + " wrong choices left.\n" + hangman_display[display] + "Keep Guessing!\n" + guessed_word + "\n")
					return
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
// TODO try to figure out how to put out a whole dictionary of dictionary
// TODO abstract handler to separate function
// TODO print letters they have used 
// TODO fix or take out ascii art