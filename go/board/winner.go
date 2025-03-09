package board

// GetAllLineCoords generates all possible winning line coordinates (rows, columns, and diagonals).
func GetAllLineCoords(gameState *GameState) [][][2]int {
	var lines [][][2]int

	// Add row coordinates
	for y := 0; y < gameState.Height; y++ {
		var row [][2]int
		for x := 0; x < gameState.Width; x++ {
			row = append(row, [2]int{y, x})
		}
		lines = append(lines, row)
	}

	// Add column coordinates
	for x := 0; x < gameState.Width; x++ {
		var col [][2]int
		for y := 0; y < gameState.Height; y++ {
			col = append(col, [2]int{y, x})
		}
		lines = append(lines, col)
	}

	// Add top-left to bottom-right diagonal
	if gameState.Width == gameState.Height { // Only for square boards
		var diag1 [][2]int
		for i := 0; i < gameState.Width; i++ {
			diag1 = append(diag1, [2]int{i, i})
		}
		lines = append(lines, diag1)
	}

	// Add top-right to bottom-left diagonal
	if gameState.Width == gameState.Height { // Only for square boards
		var diag2 [][2]int
		for i := 0; i < gameState.Width; i++ {
			diag2 = append(diag2, [2]int{i, gameState.Width - 1 - i})
		}
		lines = append(lines, diag2)
	}

	return lines
}

// GetWinner checks if any player has won based on the winning line coordinates.
func GetWinner(gameState *GameState) int {
	// Get all possible winning lines
	lines := GetAllLineCoords(gameState)

	// Iterate through each line
	for _, line := range lines {
		// Use the value of the first coordinate in the line to compare
		y, x := line[0][0], line[0][1]
		firstCell := gameState.Board[y][x]

		// Skip the line if the first cell is empty
		if firstCell == 0 {
			continue
		}

		// Check every coordinate in the line
		allMatch := true
		for _, coord := range line {
			cy, cx := coord[0], coord[1]
			if gameState.Board[cy][cx] != firstCell {
				allMatch = false
				break
			}
		}

		// If all cells in the line match, return the player value
		if allMatch {
			return firstCell
		}
	}

	// Return 0 if no winning line is found
	return 0
}

// IsBoardTie checks if the board is full and no winner exists.
func IsBoardTie(gameState *GameState) bool {
	// Iterate through the entire board
	for y := 0; y < gameState.Height; y++ {
		for x := 0; x < gameState.Width; x++ {
			// If any cell is empty, it's not a tie
			if gameState.Board[y][x] == 0 {
				return false
			}
		}
	}

	// Check if there is a winner
	if GetWinner(gameState) != 0 {
		return false // Not a tie if there is a winner
	}

	// It's a tie if the board is full and there is no winner
	return true
}
