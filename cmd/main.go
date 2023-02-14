package main

import (
	mvc "github.com/philhanna/tictactoe"
	views "github.com/philhanna/tictactoe/views"
)

func main() {
	model := mvc.NewModel()
	view := views.NewConsoleView(model)
	xPlayer := mvc.NewComputerPlayer(mvc.X)
	oPlayer := mvc.NewComputerPlayer(mvc.O)
	controller := mvc.NewController(model, view, xPlayer, oPlayer)
	controller.Run()
}
