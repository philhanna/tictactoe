package tictactoe

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// View specifies the method(s) that a view must implement.
type View interface {
	// Print displays the board.
	Print()

	// PrintMessage displays a message.
	PrintMessage(string)
}
