package main

import (
	"fmt"
	"tic-tac-toe/ai"
	"tic-tac-toe/board"
)

func main() {
	// Create a new dynamic board
	game := board.NewBoard(3, 3)
	game.RenderBoard()

	// Placeholder for player assignments (e.g., 1 = "X" and 2 = "O")
	userPlayer := 1
	aiPlayer := 2

	currentPlayer := userPlayer // User starts first

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

		// Check for a winner
		winner := board.GetWinner(game)
		if winner != 0 {
			fmt.Printf("\nPlayer %d wins!\n", winner)
			break
		}

		// Check for a tie
		if board.IsBoardTie(game) {
			fmt.Println("\nIt's a tie!")
			break
		}

		// Alternate players
		if currentPlayer == userPlayer {
			currentPlayer = aiPlayer
		} else {
			currentPlayer = userPlayer
		}
	}
}
