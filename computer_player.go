package tictactoe

import (
	"fmt"
	"math/rand"
	"time"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// ComputerPlayer is a structure with a player code (X or O)
type ComputerPlayer struct {
	code Code
	name string
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// Creates a new computer player with the code of X or O.
// Returns a pointer to this structure.
func NewComputerPlayer(code Code, name string) *ComputerPlayer {
	p := new(ComputerPlayer)
	p.code = code
	p.name = name
	return p
}

// ---------------------------------------------------------------------
// Methods (implementation of Player interface)
// ---------------------------------------------------------------------

// GetCode returns this player's code
func (self ComputerPlayer) GetCode() Code {
	return self.code
}

// GetName returns the computer player's name
func (self ComputerPlayer) GetName() string {
	return self.name
}

// GetNextMove returns the player's chosen move.
func (self ComputerPlayer) GetNextMove(c *Controller) *Location {

	time.Sleep(time.Second * 1) // Don't want the game to go too fast

	// Bonehead strategy: just pick an unoccupied square at random

	board := c.model.GetBoard()
	empties := make([]Location, 0)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == NONE {
				newLocation := Location{Row: row, Col: col}
				empties = append(empties, newLocation)
			}
		}
	}
	if len(empties) == 0 {
		panic("Board shouldn't be full; should have been checked by application")
	}

	// Choose one of the empties at random
	r := rand.Intn(len(empties))
	chosenLocation := empties[r]

	c.view.ShowMessage(fmt.Sprintf("%s chooses %s\n", self.name, chosenLocation))

	return &chosenLocation
}
