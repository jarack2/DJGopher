package games

import (
	"github.com/bwmarrin/discordgo"
	"math"
	"math/rand"
	"strconv"
)

var playerCount int = 0 // number of players in the game
var playerIds [6]string // player ids
var playerPoints [6]int // number of tokens that each player has
var playerAnswers [6]int
var playerTracker int

var currentAnswer int
var answerArray = [2]int{12,60}
var questionArray = [2]string{"How many inches are in a foot?", "How many seconds are in a minute?"} // Available questions
var currentQuestion = 0
var questionsAsked = 0
var answersCollected = 0

var trivia_game_running = false

var optedPlayers int = 0

func Trivia(s *discordgo.Session, m *discordgo.MessageCreate, trivia_game_running bool) bool {
	if m.Author.ID == s.State.User.ID { // ignoring all messages created by the bot
		return trivia_game_running
	}

	if (m.Content == "g!stop") {
		trivia_game_running = false
		resetTrivia()
		return false
	}

	if !trivia_game_running { // begin game
		trivia_game_running = true
		s.ChannelMessageSend(testing, "Lets Play Trivia!")
		s.ChannelMessageSend(testing, "How many people are playing? (can do 2 to 6 players)") // ask for player count.
	} else if (playerCount < 2 || playerCount > 6) {
		inputMessagePlayerCount(s,m)
	} else if (m.Content == "g!opt-in" && optedPlayers < playerCount) {
		playerIds[optedPlayers] = m.Author.ID
		optedPlayers++
		if (optedPlayers < playerCount) {
			s.ChannelMessageSend(testing, "Need " + strconv.Itoa(playerCount - optedPlayers) + " more players to type 'g!opt-in' to start the trivia game.")
		} else {
			s.ChannelMessageSend(testing, "That's enough players, let's begin!")
			generateQuestion(s, m)
		}
	} else {
		collectAnswer(s,m)
		answersCollected++
		if (answersCollected >= playerCount) {
			determineRoundWinner(s,m)
			answersCollected = 0
			generateQuestion(s,m)
		}
	}
	return trivia_game_running
}

func inputMessagePlayerCount(s *discordgo.Session, m *discordgo.MessageCreate) { // get the number of players
	switch m.Content { // Parse message for number of players
	case "2","two":
		playerCount = 2
	case "3","three":
		playerCount = 3
	case "4","four":
		playerCount = 4
	case "5","five":
		playerCount = 5
	case "6","six":
		playerCount = 6
	default:
		playerCount = 0
		s.ChannelMessageSend(testing, "Invalid number of players, please enter a number from 2 to 6.")
		return
	}
	s.ChannelMessageSend(testing, "You have selected " + strconv.Itoa(playerCount) + "players.")
	s.ChannelMessageSend(testing, "Type 'g!opt-in' if you would like to play trivia!")
	playerTracker = 0
	return
}

func generateQuestion(s *discordgo.Session, m *discordgo.MessageCreate) { // Select a random trivia question for users to answer
	currentQuestion = rand.Intn(2)
	s.ChannelMessageSend(testing, questionArray[currentQuestion])
	currentAnswer = answerArray[currentQuestion]
	questionsAsked++
	return
}

func collectAnswer(s *discordgo.Session, m *discordgo.MessageCreate) { // Collects the answer of an individual player
	for i := 0; i < playerCount; i++ {
		if (m.Author.ID == playerIds[i]) {
			playerAnswers[i],_ = strconv.Atoi(m.Content)
			//return (currently commented so that I can pretend to be multiple players)
			//TODO: uncomment return when more people are playing
		}
	}
	return
}

func determineRoundWinner(s *discordgo.Session, m *discordgo.MessageCreate) { // The bot reveals who had the closest answer for the current question
	var smallestDifference float64 = -1
	var topPlayer int = 0
	for i := int(0); i < playerCount; i++ {
		var answerDifference = math.Abs(float64(playerAnswers[i] - answerArray[currentQuestion]))
		if (answerDifference < smallestDifference || smallestDifference == -1) {
			smallestDifference = answerDifference
			topPlayer = i
		}
	}
	s.ChannelMessageSend(testing, "Player " + playerIds[topPlayer] + " is the winner of this round!")
}

func resetTrivia() { // reset everything
	playerCount = 0

	currentQuestion = 0
	questionsAsked = 0
	answersCollected = 0

	trivia_game_running = false

	optedPlayers = 0
}