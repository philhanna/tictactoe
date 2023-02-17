package tictactoe

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Model is an abstract representation of the game. It has no visual
// component and does not directly support user interaction.
type Model struct {
	board [][]Code
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewModel creates a model with a 3x3 empty board and returns a pointer
// to it.
func NewModel() *Model {
	board := [][]Code{
		make([]Code, 3),
		make([]Code, 3),
		make([]Code, 3),
	}
	model := Model{board}
	return &model
}

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
// If the game is not over, returns nil, nil.
func (m *Model) GetWinner() (*Code, *Vector) {
	for _, v := range Vectors {
		if code := m.CodeInAllOf(v); code != nil && *code != NONE {
			return code, v
		}
	}
	return nil, nil
}

// CodeInAllOf checks the three board locations referred to by the vector.
// If their values are all the same, returns a pointer to the code they all share.
// Otherwise, returns nil
func (m *Model) CodeInAllOf(v *Vector) *Code {
	c0 := m.board[v.loc0.Row][v.loc0.Col]
	c1 := m.board[v.loc1.Row][v.loc1.Col]
	c2 := m.board[v.loc2.Row][v.loc2.Col]
	if c1 != c0 || c2 != c0 {
		return nil
	}
	return &c0
}

// GetBoard returns a copy of the board. It is functionally read-only,
// because it does not refer back to the real board, only to its value.
func (m *Model) GetBoard() [][]Code {
	return m.board
}
