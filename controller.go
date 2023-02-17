package tictactoe

import (
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Controller struct {
	model   *Model
	view    *View
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewController is a constructor for the controller.
// It takes the following parameters:
//   - A pointer to the model
//   - A pointer to the view
// It returns a pointer to the new controller.
func NewController(model *Model, view *View) *Controller {
	c := new(Controller)
	c.model = model
	c.view = view
	return c
}
