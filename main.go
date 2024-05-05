package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Game struct {
	board      [9]string
	player     string
	turnNumber int
}

func main() {
	var game Game
	// O starts first. Can change the starter player to "O"
	game.player = "O"

	// whether game is over or not
	// initialized to false because no one has one or drawn yet
	gameOver := false
	var winner string

	for gameOver != true {
		// Print the board to show the status
		PrintBoard(game.board)

		// Asking the user to enter position
		move := AskToPlay()

		// Error Handling
		err := game.play(move)
		if err != nil {
			fmt.Println(err)
			// break loop
			continue
		}

		gameOver, winner = CheckForWinner(game.board, game.turnNumber)
	}

	// When gameOver = true
	// Print the game results

	PrintBoard(game.board)

	if winner != "" {
		fmt.Printf("\nYAY! %s wins!\n", winner)
	} else {
		fmt.Println("\nDraw!\n")
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {
	test := false
	i := 0

	// HORIZONTAL
	for i < 9 {
		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		// if no win was found in row
		// move to next row
		if !test {
			i += 3
		} else {
			// returning boolean, and who the winner is
			return true, b[i]
		}
	}

	// counter set to 0
	i = 0

	// VERTICAL
	// Testing for [0, 1, 2]
	for i < 3 {
		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		// if no win in column
		// goto next column
		if !test {
			i += 1
		} else {
			// returning boolean, and who the winner is
			return true, b[i]
		}
	}

	// PRIMARY DIAGONAL
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[i]
	}

	// SECONDARY DIAGONAL
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[i]
	}

	if n == 9 {
		return true, ""
	}
	return false, ""
}

func ClearScreen() {
	// For Windows
	cmd := exec.Command("cmd", "/c", "cls")

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		// For Linux and macOS
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		// Error Handling
	}
}

// Simple player switch function
func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return
	}
	game.player = "O"
}

// Pointer receiver method
// Game type can access play function
func (game *Game) play(pos int) error {
	// checking for ifEmpty
	if game.board[pos-1] == "" {
		// Then put the player's character here
		game.board[pos-1] = game.player

		// Switch the player
		game.SwitchPlayers()

		game.turnNumber += 1
		return nil
	}
	return errors.New("Invalid move")
}

func AskToPlay() int {
	var moveint int
	fmt.Println("Enter position [1 to 9]: ")
	// Scans users input and stores
	fmt.Scan(&moveint)
	return moveint
}

func PrintBoard(board [9]string) {
	ClearScreen()
	for i, val := range board {
		if val == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(val)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}
	}
}
