package tictactoe

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Location is a (row, col) pair of one of the squares of the board
type Location struct {
	Row int
	Col int
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of the location
func (l Location) String() string {
	return fmt.Sprintf("(%d,%d)", l.Row, l.Col)
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// ParseLocation gets a row and column number from a string. It verifies that:
//   - There are two numbers in the string
//   - Both are >= 0 and < 3
//
// If any errors are detected, returns an error message.
func ParseLocation(s string) (Location, error) {
	re := regexp.MustCompile(`-?\d+`) // Allow for negative numbers
	m := re.FindAllString(s, -1)
	if m == nil {
		errmsg := fmt.Sprintf("No numbers found in %q", s)
		err := errors.New(errmsg)
		return Location{}, err
	}
	if len(m) != 2 {
		errmsg := fmt.Sprintf("Need 2 numbers; found %d in %q", len(m), s)
		err := errors.New(errmsg)
		return Location{}, err
	}
	row, _ := strconv.Atoi(m[0])
	if row < 0 || row >= 3 {
		errmsg := fmt.Sprintf("Row must be 0, 1, or 2, not %d", row)
		err := errors.New(errmsg)
		return Location{}, err
	}
	col, _ := strconv.Atoi(m[1])
	if col < 0 || col >= 3 {
		errmsg := fmt.Sprintf("Column must be 0, 1, or 2, not %d", col)
		err := errors.New(errmsg)
		return Location{}, err
	}

	// Everything looks good
	newLoc := Location{row, col}
	return newLoc, nil
}
