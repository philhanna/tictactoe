# tictactoe
This is a sample application to illustrate the [model-view-controller
(MVC)][idMVC] pattern.

## Overview

MVC is a software architectural patterncommonly used for developing
graphical user interfaces, and features a clear [separation of
concerns][idSOC] that makes them much simpler to develop, test, and
maintain.

The application described below is a very minimal implementation of a
[Tic-Tac-Toe game][idTTT] written according to the MVC pattern. 

## Model
The model is the abstract representation of the board.  It contains only
the 3x3 matrix of squares, which may be empty, occupied by an X, or
occupied by an O.  It is created by the controller, but has no knowledge
of either the view or the controller.

The model has the following methods:
- `Get(row, col)` - Returns the value of the square at (row, col).
- `Set(row, col, value)` - Sets the square at (row, col) to X, O, or empty.
- `IsEmptySquare(row, col)` - Returns `true` if the specified square is empty. 
- `IsBoardFull()` - Returns `true` if the board is completely full.
    This indicates the game has ended in a draw
- `GetWinner()` - If the game is over, returns `X` or `O` to indicate the winner,
  or `NONE` if there is no winner. Also returns the winning vector
  (row, column, or diagonal)
- `CodeInAllOf(v Vector)` - Checks the three board locations referred to by the vector.
  If their values are all the same, returns the player code they all share.

## View
The view handles displaying the board. It is created by the controller
and has a read-only copy of the model's board.  It has the following
methods:

- `Print()` - Displays the board.
- `PrintMessage()` - Displays a message, e.g., indicating who won
or that the game was a draw, or any error messages.

The view implemented in this package is simply one that prints to the
console.  A graphical view (or a web-based one) can be substituted.  All
it needs to do is implement the same methods.

## Controller
The controller handles the action of the game. It has the single public
method:

- `Run()` - Runs the game

When created, the controller must be passed:
- The model
- The view
- The `X` and `O` players.

When `Run` is called:
- Sets the current player to `X`.
 Then in a loop, it does the following:
- Asks the model if the game is over. If so, it
  - Tells the view to announce the winner
  - Exits from the loop
Otherwise:
- Do until a legal move is received:
  - Asks the current player for his move.
  - Checks with the model to see if the move is legal
(row and column are valid, square is unoccupied).
  - If the move is illegal, tells the view to display an error message.
- Sets the value of the square in the model
- Tells the view to redisplay the board

## References
- [Github repository](https://github.com/philhanna/tictactoe)
- Wikipedia articles:
  - [ModelViewController][idMVC]
  - [Separation of concerns][idSOC]
  - [Tic-tac-toc strategy][idTTT]
- [Markdown](https://daringfireball.net/projects/markdown/syntax)
  syntax reference. Used to format this `README.md`.

[idMVC]: https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller
[idSOC]: https://en.wikipedia.org/wiki/Separation_of_concerns
[idTTT]: https://en.wikipedia.org/wiki/Tic-tac-toe
