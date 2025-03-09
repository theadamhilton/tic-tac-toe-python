package board

import (
	"testing"
)

func TestRandomAIMove(t *testing.T) {
	// Create a game state with a 3x3 board and some empty cells
	game := NewBoard(3, 3)
	game.Board[0][0] = 1 // Simulate a move on (0, 0)
	game.Board[1][1] = 2 // Simulate a move on (1, 1)

	// Call RandomAIMove and ensure the move is valid
	row, col, err := RandomAIMove(game)
	if err != nil {
		t.Fatalf("RandomAIMove returned an error: %v", err)
	}

	// Check if the move is in an empty cell
	if game.Board[row][col] != 0 {
		t.Errorf("RandomAIMove chose an occupied cell (%d, %d)", row, col)
	}

	// Ensure the move is within bounds
	if row < 0 || row >= game.Height || col < 0 || col >= game.Width {
		t.Errorf("RandomAIMove chose out-of-bounds cell (%d, %d)", row, col)
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

func TestMakeMove(t *testing.T) {
	game := NewBoard(3, 3) // Initialize a 3x3 board

	// Test valid move
	err := MakeMove(1, 1, 1, game)
	if err != nil {
		t.Errorf("MakeMove failed for a valid move: %v", err)
	}
	if game.Board[1][1] != 1 {
		t.Errorf("MakeMove did not update the board correctly at (1, 1)")
	}

	// Test out-of-bounds move
	err = MakeMove(2, 3, 3, game)
	if err == nil {
		t.Errorf("MakeMove did not return an error for an out-of-bounds move")
	}

	// Test move to an already occupied cell
	err = MakeMove(2, 1, 1, game)
	if err == nil {
		t.Errorf("MakeMove did not return an error for a move to an occupied cell")
	}

	// Test another valid move
	err = MakeMove(2, 0, 0, game)
	if err != nil {
		t.Errorf("MakeMove failed for a valid move: %v", err)
	}
	if game.Board[0][0] != 2 {
		t.Errorf("MakeMove did not update the board correctly at (0, 0)")
	}
}
