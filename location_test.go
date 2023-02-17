package tictactoe

import (
	"fmt"
	"reflect"
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

func TestParseLocation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Location
		wantErr bool
	}{
		// Good ones
		{"Good numbers", args{"1,2"}, Location{1, 2}, false},
		{"With parens", args{"(2, 2)"}, Location{2, 2}, false},

		// Bad ones
		{"Row low", args{"(-1, 2)"}, Location{}, true},
		{"Row high", args{"(3, 2)"}, Location{}, true},
		{"Col low", args{"(1, -2)"}, Location{}, true},
		{"Col high", args{"(2, 187)"}, Location{}, true},
		{"Empty string", args{""}, Location{}, true},
		{"Characters", args{"row,col"}, Location{}, true},
		{"Partial", args{"1,col"}, Location{}, true},
		{"More than 2 numbers", args{"1,2,3"}, Location{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLocation(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
