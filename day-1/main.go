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

	fmt.Printf("Max: %d", max(count(split(string(f))), 3))
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

func max(c []int, top int) int {
	sort.Slice(c, func(a, b int) bool {
		return c[b] < c[a]
	})

	t := 0
	for i, v := range c {
		if i >= top {
			break
		}
		t += v
	}

	return t
}
