package main

import (
	"fmt"
	"os"
	"strings"
)

var Opponent = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var Myself = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var BattleScore = map[int]int{
	-1: 0,
	0:  3,
	1:  6,
}

var ItemScore = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")

	t := 0
	for _, line := range lines {
		t += score(line)
	}

	fmt.Printf("Score: %d", t)

}

func score(input string) int {
	split := strings.Split(input, " ")
	opponent := Opponent[split[0]]
	myself := Myself[split[1]]

	res := battle(opponent, myself)

	return BattleScore[res] + ItemScore[myself]
}

func battle(opponent, myself string) int {
	if opponent == "Rock" {
		if myself == "Rock" {
			return 0
		}

		if myself == "Scissors" {
			return -1
		}

		if myself == "Paper" {
			return 1
		}
	}

	if opponent == "Paper" {
		if myself == "Rock" {
			return -1
		}

		if myself == "Scissors" {
			return 1
		}

		if myself == "Paper" {
			return 0
		}
	}

	if opponent == "Scissors" {
		if myself == "Rock" {
			return 1
		}

		if myself == "Scissors" {
			return 0
		}

		if myself == "Paper" {
			return -1
		}
	}

	return 0
}
