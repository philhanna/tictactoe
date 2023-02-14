package tictactoe

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Model is an abstract representation of the game. It has no visual
// component and does not directly support user interaction.
type Model struct {
	board [3][3]Code
}

// ReadOnlyModel is a set of method(s) that a read-only model
// must implement
type ReadOnlyModel interface {
	GetBoard() [3][3]Code
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewModel creates a model and returns a pointer to it.
func NewModel() *Model {
	m := new(Model)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			m.board[row][col] = NONE
		}
	}
	return m
}

/*
	How to declare a model with explicit contents:

	m := Model{[3][3]Code{
		{O, NONE, NONE},
		{NONE, NONE, X},
		{NONE, NONE, NONE},
		}}

	A board is an array of three arrays of three codes.
	An array of three codes is [3]Code{p0, p1, p2}.
	See: https://go.dev/play/p/uNH5TQ9FndX
*/

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Get returns the code value at (row, col).
func (m *Model) Get(row int, col int) Code {
	return m.board[row][col]
}

// Set sets code value at (row, col).
func (m *Model) Set(row int, col int, code Code) {
	m.board[row][col] = code
}

// IsEmptySquare returns true if the specified square is empty.
func (m *Model) IsEmptySquare(row int, col int) bool {
	code := m.board[row][col]
	return code == NONE
}

// IsBoardFull returns true if the board is completely full.
// This indicates the game has ended in a draw
func (m *Model) IsBoardFull() bool {
	nFilled := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if m.board[row][col] != NONE {
				nFilled++
			}
		}
	}
	return nFilled == 9
}

// GetWinner returns the winning player code and the winning vector.
// If the game is not over, returns NONE, NO_VECTOR
func (m *Model) GetWinner() (Code, Vector) {
	for _, v := range Vectors {
		code := m.CodeInAllOf(v)
		if code != NONE {
			return code, v
		}
	}
	return NONE, NO_VECTOR
}

// CodeInAllOf checks the three board locations referred to by the vector.
// If their values are all the same, returns the code they all share.
// Otherwise, returns NONE
func (m *Model) CodeInAllOf(v Vector) Code {
	p0 := m.board[v.loc0.Row][v.loc0.Col]
	p1 := m.board[v.loc1.Row][v.loc1.Col]
	p2 := m.board[v.loc2.Row][v.loc2.Col]
	if p1 != p0 || p2 != p0 {
		return NONE
	}
	return p0
}

// GetBoard returns a copy of the board. It is functionally read-only,
// because it does not refer back to the real board.
func (m *Model) GetBoard() [3][3]Code {
	return m.board
}
