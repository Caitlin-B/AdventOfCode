package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
)

func main() {
	inputs := utils.ScanInput(13)
	fmt.Println(day1(inputs))
	fmt.Println(day2(inputs))

}

func day1(inputs []string) int {
	tot := 0
	pattern := make([]string, 0)
	for _, in := range inputs {
		if in == "" {
			if r := getRowReflection(pattern); r != 0 {
				tot += r * 100
				pattern = []string{}
				continue
			}
			refl := swapRowCol(pattern)
			c := getRowReflection(refl)
			tot += c
			pattern = []string{}
			continue
		}
		pattern = append(pattern, in)
	}
	return tot
}

func day2(inputs []string) int {
	tot := 0
	pattern := make([]string, 0)
	for _, in := range inputs {
		if in == "" {
			if r := getRowReflectionSmudged(pattern); r != 0 {
				tot += r * 100
				pattern = []string{}
				continue
			}
			refl := swapRowCol(pattern)
			c := getRowReflectionSmudged(refl)
			tot += c
			pattern = []string{}
			continue
		}
		pattern = append(pattern, in)
	}
	return tot
}

// getRowReflection - returns 0 if no horizontal reflection
func getRowReflection(in []string) int {
	for i, r := range in {
		if i != 0 && r == in[i-1] { // if row is same as row before
			firstHalf := in[:i]
			secondHalf := in[i:]
			if isRowReflection(firstHalf, secondHalf) {
				return i
			}
		}
	}
	return 0
}

func isRowReflection(a, b []string) bool {
	a2 := make([]string, len(a))
	copy(a2, a)
	for i, j := 0, len(a2)-1; i < j; i, j = i+1, j-1 {
		a2[i], a2[j] = a2[j], a2[i]
	}
	length := len(a2)
	if len(b) < length {
		length = len(b)
	}

	for i := 0; i < length; i++ {
		if a2[i] != b[i] {
			return false
		}
	}
	return true
}

func swapRowCol(in []string) []string {
	swapped := make([]string, 0, len(in[0]))
	for i := 0; i < len(in[0]); i++ {
		col := ""
		for j := range in {
			col = col + string(in[j][i])
		}
		swapped = append(swapped, col)
	}
	return swapped
}

// isOneCharDifferent compares two rows and returns true if they are the same bar one character
func isOneCharDifferent(a, b string) bool {
	diffs := 0
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}

	return diffs == 1
}

// isRowReflectionSmudged does the same as isRowReflection but allows for one smudged row where the rows are the same in all but one character
func isRowReflectionSmudged(a, b []string, smudged bool) bool {
	a2 := make([]string, len(a))
	copy(a2, a)
	for i, j := 0, len(a2)-1; i < j; i, j = i+1, j-1 {
		a2[i], a2[j] = a2[j], a2[i]
	}
	length := len(a2)
	if len(b) < length {
		length = len(b)
	}

	for i := 1; i < length; i++ { // already know first rows are a match
		if a2[i] != b[i] {
			if !smudged && isOneCharDifferent(a2[i], b[i]) {
				smudged = true
				continue
			}
			return false
		}
	}
	return true
}

// getRowReflectionSmudged - returns 0 if no horizontal reflection
func getRowReflectionSmudged(in []string) int {
	for i, r := range in {
		if i == 0 {
			continue
		}
		oneOff := isOneCharDifferent(r, in[i-1])
		if r == in[i-1] || oneOff { // if row is same as row before
			firstHalf := in[:i]
			secondHalf := in[i:]
			if isRowReflectionSmudged(firstHalf, secondHalf, oneOff) {
				return i
			}
		}
	}
	return 0
}
