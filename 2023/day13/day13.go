package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
)

func main() {
	inputs := utils.ScanInput(13)
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

	fmt.Println(tot)
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
