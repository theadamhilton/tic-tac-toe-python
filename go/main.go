package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/ai"
	"tic-tac-toe/board"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var gameMode int
	var err error

	// Prompt the user to select a game mode
	fmt.Println("Select a game mode:")
	fmt.Println("1. User vs. AI")
	fmt.Println("2. User vs. User")
	fmt.Println("3. AI vs. AI")

	for {
		// Read user input
		input, _ := reader.ReadString('\n')

		// Trim newline or spaces from the input
		input = strings.TrimSpace(input)

		// Convert input to an integer
		gameMode, err = strconv.Atoi(input)
		if err != nil || gameMode < 1 || gameMode > 3 {
			fmt.Println("Invalid input. Please enter a number between 1 and 3.")
			continue // Retry if invalid
		}

		break
	}

	// Create a new dynamic board
	game := board.NewBoard(3, 3)
	game.RenderBoard()

	fmt.Printf("You selected game mode: %d\n", gameMode)

	// Handle game modes
	switch gameMode {
	case 1:
		playUserVsAI(game)
	case 2:
		playUserVsUser(game)
	case 3:
		playAIVsAI(game)
	}
}

func playUserVsAI(game *board.GameState) {
	// Placeholder for player assignments (e.g., 1 = "X" and 2 = "O")
	userPlayer := 1
	aiPlayer := 2
	currentPlayer := userPlayer // Selected player starts first

	// Loop to simulate a simple game with alternating turns
	for {
		// Get move based on the current player
		var row, col int
		var err error
		var found bool

		if currentPlayer == userPlayer {
			fmt.Println("Your turn! Make a move:")
			row, col, err = board.GetMove("User", game)
			if err != nil {
				fmt.Println("Error:", err)
				continue // Retry if input is invalid
			}
		} else {
			// AI's turn
			fmt.Println("Thinking...")
			row, col, found = ai.FindWinningOrBlockingMove(game, aiPlayer, userPlayer)
			if !found {
				// Fallback to a random move if no winning or blocking moves are found
				row, col, err = board.RandomAIMove(game)
				if err != nil {
					fmt.Println("Error:", err)
					break // No valid moves avilable, game ends
				}
			}
		}

		// Attempt to make the move
		err = board.MakeMove(currentPlayer, row, col, game)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Render the updated board
		game.RenderBoard()

		// Display the move made by the player
		fmt.Printf("Player %d selected column %d and row %d.\n", currentPlayer, col, row)

		// Check for winner or tie
		if checkEndGame(game, currentPlayer) {
			break
		}

		// Switch players
		currentPlayer = 3 - currentPlayer
	}
}

func playUserVsUser(game *board.GameState) {
	currentPlayer := 1

	for {
		var row, col int
		var err error

		fmt.Println("Your turn! Make a move:")
		row, col, err = board.GetMove(fmt.Sprintf("Player %d", currentPlayer), game)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		err = board.MakeMove(currentPlayer, row, col, game)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		game.RenderBoard()
		fmt.Printf("Player %d selected column %d and row %d.\n", currentPlayer, col, row)

		if checkEndGame(game, currentPlayer) {
			break
		}

		currentPlayer = 3 - currentPlayer
	}
}

func playAIVsAI(game *board.GameState) {
	aiPlayer1 := 1
	aiPlayer2 := 2
	currentPlayer := aiPlayer1

	for {
		row, col, found := ai.FindWinningOrBlockingMove(game, currentPlayer, aiPlayer2)
		var err error

		if !found {
			row, col, err = board.RandomAIMove(game)
			if err != nil {
				fmt.Printf("Player %d has no valid moves left. The game ends.\n", currentPlayer)
				break
			}
		}

		err = board.MakeMove(currentPlayer, row, col, game)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		game.RenderBoard()
		fmt.Printf("Player %d selected column %d and row %d.\n", currentPlayer, col, row)

		if checkEndGame(game, currentPlayer) {
			break
		}

		currentPlayer = 3 - currentPlayer
	}
}

func checkEndGame(game *board.GameState, _ int) bool {
	winner := board.GetWinner(game)
	if winner != 0 {
		fmt.Printf("\nPlayer %d wins!\n", winner)
		return true
	}

	if board.IsBoardTie(game) {
		fmt.Println("\nIt's a tie!")
		return true
	}

	return false
}
