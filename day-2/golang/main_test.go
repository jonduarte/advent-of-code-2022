package main

import (
	"fmt"
	"testing"
)

func Test_battle(t *testing.T) {
	tt := []struct {
		opponent string
		myself   string
		want     int
	}{
		{"Rock", "Rock", 0},
		{"Rock", "Paper", 1},
		{"Rock", "Scissors", -1},

		{"Paper", "Rock", -1},
		{"Paper", "Paper", 0},
		{"Paper", "Scissors", 1},

		{"Scissors", "Rock", 1},
		{"Scissors", "Paper", -1},
		{"Scissors", "Scissors", 0},
	}

	for _, tc := range tt {
		var result string

		switch tc.want {
		case -1:
			result = "defeat"
		case 0:
			result = "draw"
		case 1:
			result = "victory"
		}

		t.Run(fmt.Sprintf("%s x %s = %s", tc.opponent, tc.myself, result), func(t *testing.T) {
			got := battle(tc.opponent, tc.myself)
			if got != tc.want {
				t.Fatalf("expected %v, got %v", tc.want, got)
			}
		})
	}

}

func Test_score(t *testing.T) {
	tt := map[string]struct {
		input string
		want  int
	}{
		"loose": {"B X", 1},
		"draw":  {"A Y", 4},
		"win":   {"C Z", 7},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := score(tc.input)
			if got != tc.want {
				t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
			}
		})
	}
}
