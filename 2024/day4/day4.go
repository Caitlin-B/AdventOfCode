package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
)

// X = 88 M = 77 A = 65 S = 83

func main() {
	inputs := utils.ScanInput(4, 2024)

	tot := 0
	for j := range inputs {
		for i := range inputs[j] {
			//if inputs[j][i] == 88 {
			//    tot += findXmas(inputs, i, j)
			//}
			if inputs[j][i] == 65 {
				if findmasx(inputs, i, j) {
					tot++
				}
			}
		}
	}
	fmt.Print(tot)
}

func findXmas(inputs []string, i, j int) (x int) {
	mx := len(inputs[j]) - 3

	if i > 2 {
		if inputs[j][i-3:i+1] == "SAMX" {
			x++
		}
	}

	if i < mx {
		if inputs[j][i:i+4] == "XMAS" {
			x++
		}
	}

	words := make([]string, 6)
	for k := range []int{0, 1, 2, 3} {
		if j > 2 {
			words[1] += string(inputs[j-k][i])
			if i < mx {
				words[3] += string(inputs[j-k][i+k])
			}
			if i > 2 {
				words[5] += string(inputs[j-k][i-k])
			}
		}
		if j < mx {
			words[0] += string(inputs[j+k][i])
			if i < mx {
				words[2] += string(inputs[j+k][i+k])
			}
			if i > 2 {
				words[4] += string(inputs[j+k][i-k])
			}
		}
	}
	for _, w := range words {
		if w == "XMAS" {
			x++
		}
	}
	return
}

func findmasx(inputs []string, i, j int) bool {
	if i == 0 || j == 0 || i == len(inputs[j])-1 || j == len(inputs)-1 {
		return false
	}
	down := false
	up := false
	if inputs[j-1][i-1] == 77 && inputs[j+1][i+1] == 83 {
		down = true
	}
	if inputs[j-1][i-1] == 83 && inputs[j+1][i+1] == 77 {
		down = true
	}
	if inputs[j-1][i+1] == 77 && inputs[j+1][i-1] == 83 {
		up = true
	}
	if inputs[j-1][i+1] == 83 && inputs[j+1][i-1] == 77 {
		up = true
	}

	return up && down
}
