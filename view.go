package tictactoe

import (
	"fmt"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// View is a simple implementation of view that just prints to the
// console.
type View struct{}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewView creates a view and returns a pointer to it.
func NewView() *View {
	v := new(View)
	return v
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Show prints the board. This implementation uses:
//
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
func (v *View) Show(board [][]Code) {
	fmt.Printf("\n")
	fmt.Printf("    0   1   2\n")
	fmt.Printf("  +---+---+---+\n")
	fmt.Printf("0 | %s | %s | %s |\n", board[0][0], board[0][1], board[0][2])
	fmt.Printf("  +---+---+---+\n")
	fmt.Printf("1 | %s | %s | %s |\n", board[1][0], board[1][1], board[1][2])
	fmt.Printf("  +---+---+---+\n")
	fmt.Printf("2 | %s | %s | %s |\n", board[2][0], board[2][1], board[2][2])
	fmt.Printf("  +---+---+---+\n")
}

// ShowMessage displays a message
func (v *View) ShowMessage(s string) {
	fmt.Printf("%s\n", s)
}
