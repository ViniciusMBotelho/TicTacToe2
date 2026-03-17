package models

// CalculateWinner determines if there is a winner on a standard 3x3 board.
// It returns the winning CellState (PlayerX or PlayerO) or Empty if no winner exists.
func CalculateWinner(cells []CellState) CellState {
	winPatterns := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}

	for _, p := range winPatterns {
		if cells[p[0]] != Empty && cells[p[0]] != Tie &&
			cells[p[0]] == cells[p[1]] && cells[p[0]] == cells[p[2]] {
			return cells[p[0]]
		}
	}
	return Empty
}

// IsBoardFull checks if all 9 cells on the board are occupied.
func IsBoardFull(cells []CellState) bool {
	for _, c := range cells {
		if c == Empty {
			return false
		}
	}
	return true
}
