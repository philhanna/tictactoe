package tictactoe

import (
	"math/rand"
	"time"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// ComputerPlayer is a structure with a player code (X or O)
type ComputerPlayer struct {
	code Code
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// Creates a new computer player with the code of X or O.
// Returns a pointer to this structure.
func NewComputerPlayer(code Code) *ComputerPlayer {
	p := new(ComputerPlayer)
	p.code = code
	return p
}

// ---------------------------------------------------------------------
// Methods (implementation of Player interface)
// ---------------------------------------------------------------------

// GetNextMove returns the player's chosen move.
func (self ComputerPlayer) GetNextMove(board [3][3]Code) Location {

	time.Sleep(time.Second * 1) // Don't want the game to go too fase

	// Bonehead strategy: just pick an unoccupied square at random
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
		panic("Board shouldn't be full; should have been checked by controller")
	}

	// Choose one of the empties at random
	r := rand.Intn(len(empties))
	chosenLocation := empties[r]

	return chosenLocation
}

// GetCode returns this player's code
func (self ComputerPlayer) GetCode() Code {
	return self.code
}

// SetCode sets this player's code to X or O
func (self ComputerPlayer) SetCode(code Code) {
	self.code = code
}
