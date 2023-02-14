package tictactoe

// Vector is a set of three (row,col) tuplets that refer to either
//   - One of the three rows on the board
//   - One of the three columns on the board
//   - One of the two diagonals
type Vector struct {
	loc0 Location
	loc1 Location
	loc2 Location
}

// Vectors is a map of vector names to vector values.
// The names starting with "r" are rows, those with "c" are columns,
// and "d" the two diagonals.
var Vectors = map[string]Vector{
	"r0": {Location{0, 0}, Location{0, 1}, Location{0, 2}},
	"r1": {Location{1, 0}, Location{1, 1}, Location{1, 2}},
	"r2": {Location{2, 0}, Location{2, 1}, Location{2, 2}},
	"c0": {Location{0, 0}, Location{1, 0}, Location{2, 0}},
	"c1": {Location{0, 1}, Location{1, 1}, Location{2, 1}},
	"c2": {Location{0, 2}, Location{1, 2}, Location{2, 2}},
	"d0": {Location{0, 0}, Location{1, 1}, Location{2, 2}},
	"d1": {Location{0, 2}, Location{1, 1}, Location{2, 0}},
}

var NO_VECTOR = Vector{}
