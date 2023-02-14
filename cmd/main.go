package main

import (
	mvc "github.com/philhanna/tictactoe"
	views "github.com/philhanna/tictactoe/views"
	players "github.com/philhanna/tictactoe/players"
	
)

func main() {
	model := mvc.NewModel()
	view := views.NewConsoleView(model)
	xPlayer := players.NewHumanPlayer(mvc.X, "John")
	oPlayer := players.NewComputerPlayer(mvc.O)
	controller := mvc.NewController(model, view, xPlayer, oPlayer)
	controller.Run()
}
