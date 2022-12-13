package main

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	tt := map[string]struct {
		input string
		want  [][]string
	}{
		"simple": {
			"1000", [][]string{{"1000"}},
		},
		"two lines": {
			"1000\n2000", [][]string{{"1000", "2000"}},
		},
		"multiple lines simple": {
			"10\n20\n\n30", [][]string{{"10", "20"}, {"30"}},
		},
		"multiple lines": {
			`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			[][]string{
				{"1000", "2000", "3000"},
				{"4000"},
				{"5000", "6000"},
				{"7000", "8000", "9000"},
				{"10000"},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := split(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
			}
		})
	}
}

func Test_count(t *testing.T) {
	tt := map[string]struct {
		input [][]string
		want  []int
	}{
		"simple": {
			[][]string{{"10", "20"}, {"30"}}, []int{30, 30},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := count(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
			}
		})
	}
}

func Test_max(t *testing.T) {
	tt := map[string]struct {
		input []int
		want  int
		top   int
	}{
		"simple": {
			[]int{10, 30, 40, 20}, 40, 1,
		},
		"top 2": {
			[]int{10, 30, 40, 20}, 70, 2,
		},
		"edge case": {
			[]int{}, 0, 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := max(tc.input, tc.top)
			if got != tc.want {
				t.Fatalf("%s: expected, %v, got %v", name, tc.want, got)
			}
		})
	}
}
