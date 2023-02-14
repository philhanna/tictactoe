package tictactoe

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mvc "github.com/philhanna/tictactoe"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// HumanPlayer is a structure with a player code (X or O)
type HumanPlayer struct {
	code mvc.Code
	name string
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// Creates a new human player with the code of X or O.
// Returns a pointer to this structure.
func NewHumanPlayer(code mvc.Code, name string) *HumanPlayer {
	p := new(HumanPlayer)
	p.code = code
	p.name = name
	return p
}

// ---------------------------------------------------------------------
// Methods (implementation of Player interface)
// ---------------------------------------------------------------------

// GetCode returns this player's code
func (self HumanPlayer) GetCode() mvc.Code {
	return self.code
}

// GetNextMove returns the player's chosen move.
func (self HumanPlayer) GetNextMove(board [3][3]mvc.Code) mvc.Location {
	var move mvc.Location
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nEnter row, col for player %s move: ", self.code)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		move, err = mvc.ParseLocation(text)
		if err == nil {
			break
		}
		fmt.Printf("Error: %s\n", err.Error())
	}
	return move
}
