package tictactoe

import (
	"fmt"
	mvc "github.com/philhanna/tictactoe"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// ConsoleView is a simple implementation of view that just prints to
// the console.
type ConsoleView struct {
	model mvc.ReadOnlyModel
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewConsoleView creates a view and returns a pointer to it.
func NewConsoleView(model mvc.ReadOnlyModel) *ConsoleView {
	v := new(ConsoleView)
	v.model = model
	return v
}

// ---------------------------------------------------------------------
// Methods (implementation of the View interface)
// ---------------------------------------------------------------------

// Print prints the board. This implementation uses:
//
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
//	| X | X | X |
//	+---+---+---+
func (v *ConsoleView) Print() {
	myBoard := v.model.GetBoard()
	fmt.Println()
	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", myBoard[0][0], myBoard[0][1], myBoard[0][2])
	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", myBoard[1][0], myBoard[1][1], myBoard[1][2])
	fmt.Printf("+---+---+---+\n")
	fmt.Printf("| %s | %s | %s |\n", myBoard[2][0], myBoard[2][1], myBoard[2][2])
	fmt.Printf("+---+---+---+\n")
}

// PrintMessage displays a message
func (v *ConsoleView) PrintMessage(s string) {
	fmt.Printf("%s\n", s)
}
