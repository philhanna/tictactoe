package tictactoe

import (
	"fmt"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// The Controller handles the action of the game. When initialized, the controller:
//
//   - Creates the model and sets all its squares to EMPTY.
//   - Creates the view.
//   - Creates the X and O players.
//   - Sets the current player to X.
//
// Then in a loop, it does the following:
//
//   - Asks the model if the game is over. If so, it tells the view
//     to announce the winner and then exits. Otherwise:
//   - Do until a legal move is received:
//   - Asks the current player for his move.
//   - Checks with the model to see if the move is legal (square is unoccupied).
//   - If the move is illegal, tells the view to display an error message.
//   - Sets the value of the square in the model
//   - Tells the view to redisplay the board
//   - Toggles the current player
type Controller struct {
	model   *Model
	view    View
	xPlayer *Player
	oPlayer *Player
	current *Player
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewController is a constructor for the controller.
// It takes the following parameters:
//   - A pointer to the model
//   - the view
//   - the X player
//   - the O player
//
// It returns a pointer to the new controller.
func NewController(model *Model, view View, xPlayer Player, oPlayer Player) *Controller {
	c := new(Controller)
	c.model = model
	c.view = view
	c.xPlayer = &xPlayer
	c.oPlayer = &oPlayer
	c.current = &xPlayer
	return c
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Run starts the controller
func (c *Controller) Run() {
	var move Location
	var player Player
	for {

		// Ask the model if there a winner
		if player, vector := c.model.GetWinner(); player != NONE {

			// A player has won. Tell the view to announce the winning
			// player and vector
			c.view.PrintMessage(fmt.Sprintf("%s is the winner. Vector is %v", player, vector))
			break
		}

		// Ask the model if the board is full (draw)
		if c.model.IsBoardFull() {

			// The game is a draw. Tell view to announce that the game
			// is a draw and exit
			c.view.PrintMessage("Game is a draw")
			break
		}

		// Ask the current player for their move. Do this in a loop
		// until a valid move is specified.
		for {
			player = *c.current
			board := c.model.GetBoard()
			move = player.GetNextMove(board)
			if !c.model.IsEmptySquare(move.Row, move.Col) {
				errmsg := fmt.Sprintf("%s is not empty", move)
				c.view.PrintMessage(errmsg)
			}
			break
		}

		// Set the player's code in the board
		c.model.Set(move.Row, move.Col, player.GetCode())

		// Tell the view to redisplay the board
		c.view.PrintMessage(fmt.Sprintf("\nPlayer %s's move...", player.GetCode()))
		c.view.Print()

		// Toggle the players
		c.toggle()
	}
}

// toggle is an internal method that switches the current player from
// X to O, or O to X, as needed.
func (c *Controller) toggle() {
	if c.current == c.xPlayer {
		c.current = c.oPlayer
	} else {
		c.current = c.xPlayer
	}
}
