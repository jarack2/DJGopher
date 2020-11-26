package games

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type playerInfo struct {
	id          string
	name        string
	points      int
	answer      int
	hasAnswered bool
}

var playerCount int = 0 // number of players in the game
var players [6]playerInfo

var currentAnswer int
var currentQuestion = 0
var questionsAsked = 0

var trivia_game_running = false

const numberRounds = 4

var currentRound = 1

var optedPlayers int = 0

func Trivia(s *discordgo.Session, m *discordgo.MessageCreate, trivia_game_running bool) bool {
	dschannel = m.ChannelID
	if m.Author.ID == s.State.User.ID { // ignoring all messages created by the bot
		return trivia_game_running
	}

	if m.Content == "g!stop" {
		trivia_game_running = false
		resetTrivia()
		s.ChannelMessageSend(dschannel, "Trivia game stopped.")
		return false
	}

	if !trivia_game_running { // begin game
		resetTrivia()
		trivia_game_running = true
		s.ChannelMessageSend(dschannel, "Lets Play Trivia!")
		s.ChannelMessageSend(dschannel, "How many people are playing? (can do 2 to 6 players)") // ask for player count.
		optInPlayer(s, m)

	} else if playerCount < 2 || playerCount > 6 { // message has the player count
		inputMessagePlayerCount(s, m)

	} else if m.Content == "g!opt-in" && optedPlayers < playerCount { // message is a player opting in
		optInPlayer(s, m)
		if optedPlayers < playerCount {
			s.ChannelMessageSend(dschannel, "Need "+strconv.Itoa(playerCount-optedPlayers)+" more players to type 'g!opt-in' to start the trivia game.")
		} else {
			s.ChannelMessageSend(dschannel, "That's enough players, let's begin!")
			generateQuestion(s, m)
		}

	} else { // message is an answer to a trivia question
		collectAnswer(s, m)
		if canDetermineWinner() {
			determineRoundWinner(s, m)
			for i := int(0); i < playerCount; i++ {
				players[i].hasAnswered = false
			}
			currentRound++
			if !(currentRound <= numberRounds) {
				declareWinner(s, m)
				s.ChannelMessageSend(dschannel, "Stopping the game, type g!trivia to play again.")
				trivia_game_running = false
				resetTrivia()
				return false
			} else {
				generateQuestion(s, m)
			}
		}
	}
	return trivia_game_running
}

func optInPlayer(s *discordgo.Session, m *discordgo.MessageCreate) {
	players[optedPlayers].id = m.Author.ID
	players[optedPlayers].name = m.Author.Username
	optedPlayers++
}

func inputMessagePlayerCount(s *discordgo.Session, m *discordgo.MessageCreate) { // get the number of players
	switch m.Content { // Parse message for number of players
	case "2", "two":
		playerCount = 2
	case "3", "three":
		playerCount = 3
	case "4", "four":
		playerCount = 4
	case "5", "five":
		playerCount = 5
	case "6", "six":
		playerCount = 6
	default:
		playerCount = 0
		s.ChannelMessageSend(dschannel, "Invalid number of players, please enter a number from 2 to 6.")
		return
	}
	s.ChannelMessageSend(dschannel, "You have selected "+strconv.Itoa(playerCount)+" players.")
	s.ChannelMessageSend(dschannel, "Type 'g!opt-in' if you would like to play trivia!")
	return
}

func generateQuestion(s *discordgo.Session, m *discordgo.MessageCreate) { // Select a random trivia question for users to answer
	rand.Seed(time.Now().UnixNano())
	currentQuestion = rand.Intn(len(questionArray) - 1)
	s.ChannelMessageSend(dschannel, questionArray[currentQuestion].question)
	currentAnswer = questionArray[currentQuestion].answer
	questionsAsked++
	return
}

func collectAnswer(s *discordgo.Session, m *discordgo.MessageCreate) { // Collects the answer of an individual player
	for i := 0; i < playerCount; i++ {
		if m.Author.ID == players[i].id {
			answer, err := strconv.Atoi(m.Content)
			if err == nil {
				var alreadyGuessed = false
				for j := 0; j < playerCount; j++ {
					if players[j].answer == answer && players[j].hasAnswered == true && players[j].id != m.Author.ID {
						alreadyGuessed = true
					}
				}
				if !alreadyGuessed {
					players[i].answer = answer
					players[i].hasAnswered = true
				} else {
					s.ChannelMessageSend(dschannel, players[i].name+", that number was already guessed, choose a different one.")
				}
			} else {
				s.ChannelMessageSend(dschannel, players[i].name+", you can only use whole number values.")
			}
			//return (currently commented so that I can pretend to be multiple players)
			//TODO: uncomment return when more people are playing
		}
	}
	return
}

func canDetermineWinner() bool {
	for i := int(0); i < playerCount; i++ {
		if !players[i].hasAnswered {
			return false
		}
	}
	return true
}

func determineRoundWinner(s *discordgo.Session, m *discordgo.MessageCreate) { // The bot reveals who had the closest answer for the current question
	var smallestDifference float64 = -1
	var topPlayer int = 0
	for i := int(0); i < playerCount; i++ {
		var answerDifference = math.Abs(float64(players[i].answer - questionArray[currentQuestion].answer))
		if answerDifference < smallestDifference || smallestDifference == -1 {
			smallestDifference = answerDifference
			topPlayer = i
		}
	}
	players[topPlayer].points++
	s.ChannelMessageSend(dschannel, "The correct answer was: "+strconv.Itoa(questionArray[currentQuestion].answer))
	s.ChannelMessageSend(dschannel, "Player "+players[topPlayer].name+" is the winner of this round!")
}

func declareWinner(s *discordgo.Session, m *discordgo.MessageCreate) { // declares the overall winner of the match
	var topPlayer int = 0
	for i := int(0); i < playerCount; i++ {
		if players[i].points >= players[topPlayer].points {
			topPlayer = i
		}
	}
	s.ChannelMessageSend(dschannel, "Player "+players[topPlayer].name+" is the overall winner! Congrats!")
}

func resetTrivia() { // reset everything
	playerCount = 0
	players = [6]playerInfo{
		playerInfo{points: 0, name: "1", hasAnswered: false},
		playerInfo{points: 0, name: "2", hasAnswered: false},
		playerInfo{points: 0, name: "3", hasAnswered: false},
		playerInfo{points: 0, name: "4", hasAnswered: false},
		playerInfo{points: 0, name: "5", hasAnswered: false},
		playerInfo{points: 0, name: "6", hasAnswered: false},
	}

	currentQuestion = 0
	questionsAsked = 0

	currentRound = 1

	optedPlayers = 0
}
