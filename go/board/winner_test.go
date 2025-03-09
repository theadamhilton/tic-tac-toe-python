package board

import "testing"

func TestGetAllLineCoords(t *testing.T) {
	game := NewBoard(3, 3) // 3x3 board

	lines := GetAllLineCoords(game)
	expectedLineCount := 8 // 3 rows + 3 columns + 2 diagonals

	if len(lines) != expectedLineCount {
		t.Errorf("Expected %d lines, but got %d", expectedLineCount, len(lines))
	}
}

func TestGetWinner(t *testing.T) {
	game := NewBoard(3, 3)

	// Test no winner
	winner := GetWinner(game)
	if winner != 0 {
		t.Errorf("Expected no winner, but got player %d", winner)
	}

	// Test winner in a row
	game.Board[0][0], game.Board[0][1], game.Board[0][2] = 1, 1, 1
	winner = GetWinner(game)
	if winner != 1 {
		t.Errorf("Expected winner to be player 1, but got player %d", winner)
	}

	// Test winner in a column
	game = NewBoard(3, 3)
	game.Board[0][0], game.Board[1][0], game.Board[2][0] = 2, 2, 2
	winner = GetWinner(game)
	if winner != 2 {
		t.Errorf("Expected winner to be player 2, but got player %d", winner)
	}

	// Test winner in a diagonal
	game = NewBoard(3, 3)
	game.Board[0][0], game.Board[1][1], game.Board[2][2] = 1, 1, 1
	winner = GetWinner(game)
	if winner != 1 {
		t.Errorf("Expected winner to be player 1, but got player %d", winner)
	}

	// Test no winner with a mix of cells
	game = NewBoard(3, 3)
	game.Board[0][0], game.Board[0][1], game.Board[0][2] = 1, 2, 1
	game.Board[1][0], game.Board[1][1], game.Board[1][2] = 2, 1, 2
	game.Board[2][0], game.Board[2][1], game.Board[2][2] = 2, 1, 2

	winner = GetWinner(game)
	isTie := IsBoardTie(game)

	if winner != 0 {
		t.Errorf("Expected no winner, but got player %d", winner)
	}

	if !isTie {
		t.Errorf("Expected a tie, but IsBoardTie returned false")
	}
}
