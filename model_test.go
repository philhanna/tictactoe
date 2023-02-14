package tictactoe

import (
	"testing"
)

func TestModel_IsEmptySquare(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name  string
		model Model
		args  args
		want  bool
	}{
		{"All empty", Model{[3][3]Code{}}, args{1, 1}, true},
		{"Not all empty", Model{[3][3]Code{
			{O, NONE, NONE},
			{NONE, NONE, X},
			{NONE, NONE, NONE},
		}}, args{1, 2}, false},
		{"Not right location", Model{[3][3]Code{
			{O, NONE, NONE},
			{NONE, NONE, X},
			{NONE, NONE, NONE},
		}}, args{2, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.model
			if got := m.IsEmptySquare(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("Model.IsEmptySquare(%d,%d) = %v, want %v", tt.args.row, tt.args.col, got, tt.want)
			}
		})
	}
}

func TestModel_IsBoardFull(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  bool
	}{
		{"All empty", Model{[3][3]Code{}}, false},
		{"No winner", Model{[3][3]Code{
			{O, NONE, NONE},
			{NONE, NONE, X},
			{NONE, NONE, NONE},
		}}, false},
		{"X winner", Model{[3][3]Code{
			{X, NONE, NONE},
			{NONE, X, O},
			{NONE, NONE, X},
		}}, false},
		{"Full", Model{[3][3]Code{
			{X, O, X},
			{O, X, O},
			{X, X, O},
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.model
			if got := m.IsBoardFull(); got != tt.want {
				t.Errorf("Model.IsBoardFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
