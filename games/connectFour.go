package games

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var player1 = ""
var player2 = ""
var activePlayer = ""
var playersFull bool = false
var lastPieceX int
var lastPieceY int
var gameWin bool = false

const ROWS = 6
const COLS = 6

var xChoice []string = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣"}
var formatBoard [ROWS][COLS]int
var blankBoard [ROWS][COLS]int
var emptyPiece = "⚪"
var p1Piece = "🔴"
var p2Piece = "🔵"
var boardMessage = ""
var winner string
var playerTurn string

//ConnectFour driver
func ConnectFour(s *discordgo.Session, m *discordgo.MessageCreate, connectFourRunning bool, playerStart string) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	dschannel = m.ChannelID

	activePlayer = playerStart
	if !connectFourRunning {
		activePlayer = playerStart
		player1 = playerStart
		playerTurn = player1
		s.ChannelMessageSend(dschannel, "Lets Play ConnectFour!")
		//	playerJoin(s, m) //loops until player 2 joins
		boardToString() //string representation of board
		s.ChannelMessageSend(dschannel, boardMessage)
		s.ChannelMessageSend(dschannel, "Player2 opt in with g!gameJoin")

	} else {
		if !gameWin {
			if !playersFull {
				playerJoin(s, m)
			} else {
				if (playerStart == player1 || playerStart == player2) && (playerTurn == playerStart) {
					connectFourRunning = boardFull()
					dropPiece(s, m, player1, player2)
					if gameWin {
						return
					}
					boardToString()

					s.ChannelMessageSend(dschannel, "Player: "+activePlayer+" turn")
					s.ChannelMessageSend(dschannel, boardMessage)
					if activePlayer == player1 {
						s.ChannelMessageSend(dschannel, "Switching to: "+player2+" turn")
						playerTurn = player2
					} else if activePlayer == player2 {
						s.ChannelMessageSend(dschannel, "Switching to: "+player1+" turn")
						playerTurn = player1
					}
				}
			}
		} else {
			if playerStart != player1 || playerStart != player2 && playerTurn == playerStart {
				s.ChannelMessageSend(dschannel, "Game Won by: "+ activePlayer)
				s.ChannelMessageSend(dschannel, "Restart using g!connect4 stop")
				return
			}

		}

	}
	return
}

func checkWin(lastValue int) bool {

	for j := 0; j < 3; j++ {
		for i := 0; i < 6; i++ {
			if formatBoard[i][j] == lastValue && formatBoard[i][j+1] == lastValue && formatBoard[i][j+2] == lastValue && formatBoard[i][j+3] == lastValue {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < 3; i++ {
		for j := 0; j < 6; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i+1][j] == lastValue && formatBoard[i+2][j] == lastValue && formatBoard[i+3][j] == lastValue {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i-1][j+1] == lastValue && formatBoard[i-2][j+2] == lastValue && formatBoard[i-3][j+3] == lastValue {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			if formatBoard[i][j] == lastValue && formatBoard[i-1][j-1] == lastValue && formatBoard[i-2][j-2] == lastValue && formatBoard[i-3][j-3] == lastValue {
				return true
			}
		}
	}
	return false
}

func setActive(activePlayer string) {
	if activePlayer == player1 {
		activePlayer = player2
	} else if activePlayer == player2 {
		activePlayer = player1
	}
}

func boardFull() bool {
	for i := 0; i < ROWS-1; i++ {
		for j := 0; j < COLS-1; j++ {
			if formatBoard[i][j] != 0 {
				return false
			}
		}
	}
	return true

}

func checkSpace(input int, pieceVal int, s *discordgo.Session) bool {
	i := ROWS-1
	var emptySpace bool = false
	for i > -1 {
		if formatBoard[i][input] != 0 { //checks the input column, row by row
			i--
			if i < 0 {
				return emptySpace //false if no empty pieces in column
			}
		} else {
			formatBoard[i][input] = pieceVal       //sets empty piece to activeplayer piece
			gameWin = checkWin(pieceVal) //checks to see if game is over
			if gameWin {
				emptySpace = true
				return emptySpace
			}
			emptySpace = true
			break
		}
	}
	return emptySpace
}

func dropPiece(s *discordgo.Session, m *discordgo.MessageCreate, player1 string, player2 string) {
	if m.Author.Username != activePlayer {
		s.ChannelMessageSend(dschannel, "Error: You are not the active Player!")
	} else {
		input, err := strconv.Atoi(m.Content)
		input--
		if err != nil {
			s.ChannelMessageSend(dschannel, "Error: input not a number")
		}
		if input < 0 || input > COLS {
			s.ChannelMessageSend(dschannel, "Error: input must be in range 0 to "+strconv.Itoa(COLS))
			return
		}

		var pieceVal int = 0
		if activePlayer == player1 {
			pieceVal = 1
		} else {
			pieceVal = 2
		}
		//input to change piece on board to activePlayer color
		check := checkSpace(input, pieceVal, s)
		if gameWin {
			s.ChannelMessageSend(dschannel, "Game Won by: "+activePlayer)
			s.ChannelMessageSend(dschannel, "Restart using g!connect4 stop")
			return
		}
		if !check {
			s.ChannelMessageSend(dschannel, "Error: Column Full input another column")
			s.ChannelMessageSend(dschannel, "If entire board if full restart using g!connect4 stop")
		} else {

		}
	}

}

func playerJoin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "g!gameJoin" {
		player2 = m.Author.Username
		playersFull = true
		s.ChannelMessageSend(dschannel, "Added player 2: "+player2)
		s.ChannelMessageSend(dschannel, "Player " + player1 + " turn")
		s.ChannelMessageSend(dschannel, "Player " + player1 + " turn")
		return
	}
	if m.Content != "g!gameJoin" {
		s.ChannelMessageSend(dschannel, "Error No Player 2")
		s.ChannelMessageSend(dschannel, "Exit with g!connect4 stop")
	}

}

func boardToString() {
	boardMessage = ""
	var piece int

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			piece = formatBoard[i][j]
			switch piece {
			case 0:
				boardMessage += emptyPiece
			case 1:
				boardMessage += p1Piece
			case 2:
				boardMessage += p2Piece
			}
		}
		boardMessage += "\n"
	}
	boardMessage += "1️⃣" + "2️⃣" + "3️⃣" + "4️⃣" + "5️⃣" + "6️⃣"
	boardMessage += "\n"
}

func ConnectFourReset(run bool) {
	run = false
	formatBoard = blankBoard
	gameWin = false
	player2 = ""
	playersFull = false
}