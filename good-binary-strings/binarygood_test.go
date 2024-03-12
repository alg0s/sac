package binarygood

import (
	"testing"
)

func TestLargestGood(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", "", ""},
		{"invalid string", "abc", ""},
		{"good string 1", "11011000", "11100100"},
		{"good string 2", "1100", "1100"},
		{"good string 3", "1101001100", "1101001100"},
		{"good string 4", "0110", ""},
		{"bad string 1", "11011001", ""},
		{"good string 5", "10", "10"},
		{"bad string 6", "101", ""},
		{"bad string 7", "01", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := largestGood(test.input)
			if got != test.want {
				t.Errorf("largestGood(%s) = %s; want %s", test.input, got, test.want)
			}
		})
	}
}
