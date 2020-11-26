// TODO: allow different users to play differet games
// TODO: make sure word cannot be whitespace

package games

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

var dictionary []string    // a makeshift dictionary
var chosen_word string     // the word picked from the dictionary
var guessed_letters string // the letters that the user has guessed
var guessed_word string    // the word that the user is currently dschannel
var dschannel string

var movesLeft = 7 // how many wrong moves the user has left
var display = 0   // where the display is in the array
func Hangman(s *discordgo.Session, m *discordgo.MessageCreate, game_running bool) {
	dschannel := m.ChannelID
	if !game_running {
		createWordBank()
		chosen_word = dictionary[rand.Intn(len(dictionary))]
		fmt.Println(chosen_word)
		s.ChannelMessageSend(dschannel, "Lets Play Hangman!")
		s.ChannelMessageSend(dschannel, Hangman_display[display])

		for i := 0; i < len(chosen_word); i++ { // creates the length of the string in underscores to print for hangman
			if chosen_word[i] == ' ' {
				guessed_word += " "
			} else {
				guessed_word += "-"
			}
		}

		s.ChannelMessageSend(dschannel, guessed_word)
	} else {
		inputMessage(s, m)
	}
	return
}

func Restart(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(dschannel, "Restarting...")
	movesLeft = 7
	display = 0
	guessed_word = ""
	guessed_letters = ""
	return
}

func createWordBank() {
	file, err := os.Open("./games/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			dictionary = append(dictionary, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// callback function for when the user guesses a letter
func inputMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID { // ignoring all messages created by the bot
		return
	}

	if lost() || won() {
		return
	}

	user_input := m.Content

	if len(user_input) > 1 || !unicode.IsLetter(rune(user_input[0])) { // the user input should only be a single character and should be a letter [a-z]
		s.ChannelMessageSend(dschannel, "The user input is invalid. Please try again.\n")
		return
	}

	if strings.Contains(strings.ToLower(guessed_letters), strings.ToLower(user_input)) { // if the user has already guessed a letter
		s.ChannelMessageSend(dschannel, "You have already guessed that letter. Please try another one.\n")
		return
	} else {
		guessed_letters += strings.ToLower(user_input) + " "

		if strings.Contains(strings.ToLower(chosen_word), strings.ToLower(user_input)) { // check to see if the users input matches the word
			replaceWordWithSuccessfulGuess(user_input)
			s.ChannelMessageSend(dschannel, guessed_word)
			if won() {
				s.ChannelMessageSend(dschannel, "You guessed the word!\n"+"You can play again with the command: g!hangman stop\n")
				return
			}
			return
		} else { // chosen letter is not in the word
			movesLeft--

			if lost() { // the user has lost the game
				s.ChannelMessageSend(dschannel, "You failed to guess: \""+chosen_word+"\". Better luck next time.\n")
				s.ChannelMessageSend(dschannel, "You can play again with the command: g!hangman stop\n")
				return
			} else { // wrong input, but the user hasnt lost yet
				display++
				s.ChannelMessageSend(dschannel, user_input+" is not in the word.\n"+"You have "+strconv.Itoa(movesLeft)+" wrong choices left.\n"+Hangman_display[display]+"\nKeep Guessing!\n"+guessed_word+"\n")
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

func lost() bool {
	if movesLeft == 0 {
		return true
	}
	return false
}

func won() bool {
	if chosen_word == guessed_word {
		return true
	}
	return false
}
