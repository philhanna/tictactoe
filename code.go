package tictactoe

// Code is an X or O code in a single cell in a tic-tac-toe board model.
type Code byte

const (
	NONE Code = iota
	X
	O
)

// String returns a string representation of the player (X or O)
func (s Code) String() string {
	switch s {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}
