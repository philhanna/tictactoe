package tictactoe

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Player specifies the method(s) that either a human or computer
// player must implement.
type Player interface {
	// GetCode returns this player's code (X or O)
	GetCode() Code

	// GetName returns this player's name
	GetName() string
	
	// GetNextMove returns the row and column of the next move
	GetNextMove(c *Controller) *Location
}
