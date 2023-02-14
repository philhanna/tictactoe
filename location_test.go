package tictactoe

import (
	"fmt"
	"testing"
)

func TestLocationString(t *testing.T) {
	location := Location{Row: 1, Col: 2}
	have := fmt.Sprintf("%s", location)
	want := "(1,2)"
	if have != want {
		t.Errorf("have=%s,want=%s", have, want)
	}
}
