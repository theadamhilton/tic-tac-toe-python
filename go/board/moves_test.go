package board

import (
	"testing"
)

func TestRandomAIMove(t *testing.T) {
	// Create a game state with a 3x3 board and some empty cells
	game := NewBoard(3, 3)
	game.Board[0][0] = 1 // Simulate a move on (0, 0)
	game.Board[1][1] = 2 // Simulate a move on (1, 1)

	// Call randomAIMove and ensure the move is valid
	row, col, err := randomAIMove(game)
	if err != nil {
		t.Fatalf("randomAIMove returned an error: %v", err)
	}

	// Check if the move is in an empty cell
	if game.Board[row][col] != 0 {
		t.Errorf("randomAIMove chose an occupied cell (%d, %d)", row, col)
	}

	// Ensure the move is within bounds
	if row < 0 || row >= game.Height || col < 0 || col >= game.Width {
		t.Errorf("randomAIMove chose out-of-bounds cell (%d, %d)", row, col)
	}
}

func TestGetMoveAI(t *testing.T) {
	// Create a game state with a 3x3 board
	game := NewBoard(3, 3)

	// Simulate a call to GetMove for the AI
	row, col, err := GetMove("AI", game)
	if err != nil {
		t.Fatalf("GetMove returned an error for AI: %v", err)
	}

	// Check if the move is valid
	if game.Board[row][col] != 0 {
		t.Errorf("GetMove (AI) chose an occupied cell (%d, %d)", row, col)
	}
}
