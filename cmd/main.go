package main

import (
	"fmt"

	ttt "github.com/philhanna/tictactoe"
)

// The application handles the action of the game.
//
//   - Creates the X and O players
//   - Sets the current player to X
//   - Creates the model and sets all its squares to EMPTY
//   - Creates the view
//   - Creates the controller
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
func main() {

	// Create the X and O players
	xPlayer := ttt.NewHumanPlayer(ttt.X, "John")
	oPlayer := ttt.NewComputerPlayer(ttt.O, "Computer")

	// Set the current player to X
	var player ttt.Player
	player = xPlayer

	// Create the model, view, and controller
	m := ttt.NewModel()
	v := ttt.NewView()
	c := ttt.NewController(m, v)

	// Play the game until there is a winner or the board is full

outer:
	for {

		// Ask the model if there a winner
		if winner, vector := m.GetWinner(); winner != nil {

			// A player has won. Tell the view to announce the winning
			// player and vector
			v.Print(m.GetBoard())
			msg := fmt.Sprintf("%s is the winner. Vector is %v", winner, vector)
			v.PrintMessage(msg)
			break
		}

		// Ask the model if the board is full (draw)
		if m.IsBoardFull() {

			// The game is a draw. Tell view to announce that the game
			// is a draw and exit
			v.PrintMessage("Game is a draw")
			break
		}

		// Ask the current player for their move. Do this in a loop
		// until a valid move is specified.
		for {

			// Tell the view to redisplay the board
			msg := fmt.Sprintf("\n%s's move...", player.GetName())
			v.PrintMessage(msg)
			v.Print(m.GetBoard())

			// Ask the player for their next location
			location := player.GetNextMove(c)
			if location == nil {
				break outer
			}
			if !m.IsEmptySquare(location.Row, location.Col) {
				errmsg := fmt.Sprintf("%s is not empty", location)
				v.PrintMessage(errmsg)
				continue // Try again
			}

			// Set the player's code in the board
			m.Set(location.Row, location.Col, player.GetCode())

			// Toggle the players
			if player == xPlayer {
				player = oPlayer
			} else {
				player = xPlayer
			}
			break
		}
	}
}
