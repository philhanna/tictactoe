package tictactoe

import (
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Vector is a set of three (row,col) tuplets that refer to either
//   - One of the three rows on the board
//   - One of the three columns on the board
//   - One of the two diagonals
type Vector struct {
	loc0 Location
	loc1 Location
	loc2 Location
}

// Vectors is a slice of vector pointers for each of the rows, columns, and diagonals.
var Vectors = []*Vector{
	{Location{0, 0}, Location{0, 1}, Location{0, 2}},
	{Location{1, 0}, Location{1, 1}, Location{1, 2}},
	{Location{2, 0}, Location{2, 1}, Location{2, 2}},
	{Location{0, 0}, Location{1, 0}, Location{2, 0}},
	{Location{0, 1}, Location{1, 1}, Location{2, 1}},
	{Location{0, 2}, Location{1, 2}, Location{2, 2}},
	{Location{0, 0}, Location{1, 1}, Location{2, 2}},
	{Location{0, 2}, Location{1, 1}, Location{2, 0}},
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of the vector
func (v *Vector) String() string {
	return strings.Join([]string{
		v.loc0.String(),
		v.loc1.String(),
		v.loc2.String(),
	}, ", ")
}
