package tictactoe

import "fmt"

// Location is a (row, col) pair of one of the squares of the board
type Location struct {
	Row int
	Col int
}

// String returns a string representation of the location
func (l Location) String() string {
	return fmt.Sprintf("(%d,%d)", l.Row, l.Col)
}
