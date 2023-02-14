package tictactoe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nEnter row, col for player %s move: ", self.code)
		text, _ := reader.ReadString('\n')
		re := regexp.MustCompile(`\d+`)
		result := re.FindAllString(text, -1)
		if result == nil || len(result) != 2 {
			fmt.Printf("Error: must enter two numbers.\n")
		} else {
			rowString, colString := result[0], result[1]
			row, _ := strconv.Atoi(rowString)
			col, _ := strconv.Atoi(colString)
			if row < 0 || row >= 3 {
				fmt.Printf("Error: row is out of bounds\n")
			} else {
				if col < 0 || col >= 3 {
					fmt.Printf("Error: col is out of bounds\n")
				} else {
					move = mvc.Location{Row: row, Col: col}
					break
				}
			}
		}
	}
	return move
}
