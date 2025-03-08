package board

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Create a custom random number generator
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// getMove fetches the move from either a human or an AI player.
func GetMove(player string, gameState *GameState) (int, int, error) {
	if player == "AI" {
		// Call AI logic to generate a move
		return randomAIMove(gameState)
	} else {
		// Prompt the human player for their move
		return userMoveInput(gameState)
	}
}

// userMoveInput handles move input for user players.
func userMoveInput(gameState *GameState) (int, int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt for X coordinate (Column)
		fmt.Print("Enter the column (0 to ", gameState.Width-1, "): ")
		scanner.Scan()
		xInput := scanner.Text()
		x, err := strconv.Atoi(xInput)
		if err != nil || x < 0 || x >= gameState.Width {
			fmt.Println("Invalid column input. Please try again.")
			continue
		}

		// Prompt for Y coordinate (Row)
		fmt.Print("Enter the row (0 to ", gameState.Height-1, "): ")
		scanner.Scan()
		yInput := scanner.Text()
		y, err := strconv.Atoi(yInput)
		if err != nil || y < 0 || y >= gameState.Height {
			fmt.Println("Invalid row input. Please try again.")
			continue
		}

		// Check if the cell is already occupied
		if gameState.Board[y][x] != 0 {
			fmt.Printf("Cell (%d, %d) is already occupied. Please try again.\n", x, y)
			continue
		}

		return y, x, nil // Valid move
	}
}

// randomAIMove generates a random move for the AI player.
func randomAIMove(gameState *GameState) (int, int, error) {
	var emptyCells []struct{ Row, Col int }
	for y := 0; y < gameState.Height; y++ {
		for x := 0; x < gameState.Width; x++ {
			if gameState.Board[y][x] == 0 {
				emptyCells = append(emptyCells, struct{ Row, Col int }{Row: y, Col: x})
			}
		}
	}

	if len(emptyCells) == 0 {
		return -1, -1, fmt.Errorf("no valid moves available")
	}

	// Pick a random empty cell using the custom RNG
	choice := emptyCells[rng.Intn(len(emptyCells))]
	return choice.Row, choice.Col, nil
}
