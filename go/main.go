package main

import (
	"fmt"
	"tic-tac-toe/board"
)

func main() {
	// Create a new dynamic board
	game := board.NewBoard(3, 3)
	game.RenderBoard()

	// Placeholder for determining player type (e.g., user vs AI)
	currentPlayer := "User"

	// Get a move from the current player
	if currentPlayer == "User" {
		fmt.Println("\nYour turn! Make a move:")
	}
	row, col, err := board.GetMove(currentPlayer, game)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return // Exit if an error occurs
	}

	// Update the board with the move (temporarily assume Player 1, "X")
	game.Board[row][col] = 1

	// Render the board again
	game.RenderBoard()

	// Display the move made by the player
	fmt.Printf("You selected column %d and row %d.\n", col, row)

	// http.HandleFunc("/makeMove", makeMove)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
