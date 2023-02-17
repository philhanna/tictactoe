package tictactoe

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// HumanPlayer is a structure with a player code (X or O)
type HumanPlayer struct {
	code Code
	name string
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// Creates a new human player with the code of X or O.
// Returns a pointer to this structure.
func NewHumanPlayer(code Code, name string) *HumanPlayer {
	p := new(HumanPlayer)
	p.code = code
	p.name = name
	return p
}

// ---------------------------------------------------------------------
// Methods (implementation of Player interface)
// ---------------------------------------------------------------------

// GetCode returns this player's code
func (self HumanPlayer) GetCode() Code {
	return self.code
}

// GetName returns the human player's name
func (self HumanPlayer) GetName() string {
	return self.name
}

// GetNextMove returns the player's chosen move.
func (self HumanPlayer) GetNextMove(c *Controller) *Location {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Ask player for a (row,col) pair
		msg := fmt.Sprintf("\nEnter row, col for %s move:", self.code)
		c.view.PrintMessage(msg)

		// Read the player's response
		scanner.Scan()
		text := scanner.Text()
		if text[0] == 'q' {
			return nil
		}

		// Parse the response
		location, err := ParseLocation(text)
		validMove := err == nil
		if validMove {

			// If the location is valid, return it
			return &location
			
		} else {

			// Otherwise, display error message and prompt again
			msg := fmt.Sprintf("Error: %s is not a valid location", text)
			c.view.PrintMessage(msg)
		}
	}
}
