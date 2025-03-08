package board

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewBoard(t *testing.T) {
	// TestNewBoard ensures that the board initializes with the
	// correct dimensions and empty cells.
	width, height := 3, 3
	game := NewBoard(width, height)

	// Check dimensions
	if len(game.Board) != height {
		t.Errorf("Expected height %d, got %d", height, len(game.Board))
	}
	for i := 0; i < height; i++ {
		if len(game.Board[i]) != width {
			t.Errorf("Expected width %d in row %d, got %d", width, i, len(game.Board[i]))
		}
	}

	// Check that all cells are initialized to 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if game.Board[y][x] != 0 {
				t.Errorf("Expected cell (%d, %d) to be 0, got %d", y, x, game.Board[y][x])
			}
		}
	}
}

// TestRenderBoard ensures that RenderBoard outputs correctly
// for an empty board.
func TestRenderBoard(t *testing.T) {
	width, height := 3, 3
	game := NewBoard(width, height)

	// Capture the output of RenderBoard
	output := captureRenderOutput(game.RenderBoard)

	// Validate output for an empty board
	expectedHeader := "   0 1 2" // Adjust dynamically if needed
	if !strings.Contains(output, expectedHeader) {
		t.Errorf("Expected header '%s', but not found in output\n%s", expectedHeader, output)
	}

	expectedRows := []string{
		"0 | . . . |",
		"1 | . . . |",
		"2 | . . . |",
	}
	for _, expectedRow := range expectedRows {
		if !strings.Contains(output, expectedRow) {
			t.Errorf("Expected row '%s', but not found in output\n%s", expectedRow, output)
		}
	}
}

// Helper function to capture RenderBoard's output
func captureRenderOutput(renderFunc func()) string {
	// Use a buffer to capture the output
	var output bytes.Buffer

	// Save the original standard output
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore stdout after function execution

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the render function
	renderFunc()

	// Close the writer and capture the output
	w.Close()
	output.ReadFrom(r)

	return output.String()
}
