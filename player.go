package tictactoe

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Player specifies the method(s) that either a human or computer
// player must implement.
type Player interface {
	// GetCode returns this player's code (X or O)
	GetCode() Code

	// SetCode sets this player's code to X or O
	SetCode(code Code)

	// GetNextMove returns the
	GetNextMove(board [3][3]Code) Location
}
