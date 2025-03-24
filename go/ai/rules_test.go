package ai

import (
	"testing"
	"tic-tac-toe/board" // Import the board package for GameState and related functions
)

func TestFindWinningOrBlockingMove(t *testing.T) {
	// Test Case 1: AI has a winning move
	game := board.NewBoard(3, 3)
	game.Board[0][0], game.Board[0][1] = 2, 2 // AI player
	game.Board[0][2] = 0                      // Winning move for AI
	y, x, found := FindWinningOrBlockingMove(game, 2, 1)
	if !found || y != 0 || x != 2 {
		t.Errorf("Expected winning move at (0, 2), but got (%d, %d, %t)", y, x, found)
	}

	// Test Case 2: Opponent has a winning move, AI blocks it
	game = board.NewBoard(3, 3)
	game.Board[1][0], game.Board[1][1] = 1, 1 // Opponent player
	game.Board[1][2] = 0                      // Blocking move needed
	y, x, found = FindWinningOrBlockingMove(game, 2, 1)
	if !found || y != 1 || x != 2 {
		t.Errorf("Expected blocking move at (1, 2), but got (%d, %d, %t)", y, x, found)
	}

	// Test Case 3: No winning or blocking move
	game = board.NewBoard(3, 3)
	game.Board[0][0], game.Board[0][1], game.Board[0][2] = 1, 2, 1
	game.Board[1][0], game.Board[1][1], game.Board[1][2] = 2, 1, 2
	game.Board[2][0], game.Board[2][1], game.Board[2][2] = 2, 1, 2 // Full board, no winning moves
	y, x, found = FindWinningOrBlockingMove(game, 2, 1)
	if found {
		t.Errorf("Expected no moves, but got (%d, %d, %t)", y, x, found)
	}
}

// Helper to create a GameState with predefined board values
func createGameState(state [][]int) *board.GameState {
	game := board.NewBoard(3, 3)
	for i, row := range state {
		copy(game.Board[i], row) // Use copy instead of a loop
	}
	return game
}

// Test the Evaluate function
func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		boardState [][]int
		expected   int
	}{
		{"AI wins", [][]int{
			{2, 2, 2},
			{1, 0, 0},
			{1, 0, 0},
		}, 10},
		{"Human wins", [][]int{
			{1, 1, 1},
			{2, 0, 0},
			{2, 0, 0},
		}, -10},
		{"Tie", [][]int{
			{1, 2, 1},
			{2, 1, 2},
			{2, 1, 2},
		}, 0},
		{"Ongoing", [][]int{
			{1, 0, 1},
			{2, 1, 2},
			{2, 0, 2},
		}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := createGameState(test.boardState)
			result := evaluate(game, 2, 1)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Test the FindBestMove function
func TestFindBestMove(t *testing.T) {
	tests := []struct {
		name       string
		boardState [][]int
		expected   [2]int // Expected row and column
	}{
		{"AI wins", [][]int{
			{2, 2, 0},
			{1, 0, 0},
			{1, 0, 0},
		}, [2]int{0, 2}},
		{"Human block", [][]int{
			{1, 1, 0},
			{2, 2, 0},
			{0, 0, 0},
		}, [2]int{0, 2}},
		{"Choose center", [][]int{
			{1, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}, [2]int{1, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := createGameState(test.boardState)
			row, col := FindBestMove(game, 2, 1)
			if [2]int{row, col} != test.expected {
				t.Errorf("expected move %v, got [%d, %d]", test.expected, row, col)
			}
		})
	}
}

// Test edge cases, such as full board or one empty cell
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		boardState [][]int
		expected   [2]int
	}{
		{"Full board", [][]int{
			{1, 2, 1},
			{2, 1, 2},
			{2, 1, 2},
		}, [2]int{-1, -1}}, // No valid move
		{"One empty cell", [][]int{
			{1, 2, 1},
			{2, 0, 2},
			{2, 1, 1},
		}, [2]int{1, 1}}, // Only valid move
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := createGameState(test.boardState)
			row, col := FindBestMove(game, 2, 1)
			if [2]int{row, col} != test.expected {
				t.Errorf("expected move %v, got [%d, %d]", test.expected, row, col)
			}
		})
	}
}
