package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input-2")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Max: %d", max(count(split(string(f)))))
}

func split(t string) [][]string {
	var groups [][]string

	elves := strings.Split(t, "\n\n")
	for _, line := range elves {
		groups = append(groups, strings.Split(line, "\n"))
	}

	return groups
}

func count(c [][]string) []int {
	var res []int

	for _, group := range c {
		t := 0
		for _, line := range group {
			v, err := strconv.Atoi(line)
			if err != nil {
				v = 0
			}

			t += v
		}
		res = append(res, t)
	}

	return res
}

func max(c []int) int {
	sort.Ints(c)

	if len(c) > 0 {
		return c[len(c)-1]
	} else {
		return 0
	}
}
