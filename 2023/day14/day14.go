package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strings"
)

func main() {
	inputs := utils.ScanInput(14)
	//inputs = rollNorth(inputs, 2, 1)

	for i := 1; i < len(inputs); i++ { // skip first row, nothing will roll from there
		for j := 0; j < len(inputs[i]); j++ {
			if string(inputs[i][j]) == "O" {
				inputs = rollNorth(inputs, j, i)
			}
		}
	}

	rows := len(inputs)
	tot := 0
	for i, in := range inputs {
		mult := rows - i
		tot += strings.Count(in, "O") * mult
	}

	fmt.Print(tot)
}

// rollNorth takes a list of inputs and the position of a rock and moves the rock as far north as possible,
// returning the inputs with the rock in its new position
func rollNorth(inputs []string, x, y int) []string {
	// remove rock from old position
	inputs[y] = inputs[y][:x] + "." + inputs[y][x+1:]

	// move rock to new position
	for i := y - 1; i >= 0; i-- {
		if string(inputs[i][x]) == "#" || string(inputs[i][x]) == "O" {
			inputs[i+1] = inputs[i+1][:x] + "O" + inputs[i+1][x+1:]
			break
		}
		if i == 0 {
			inputs[i] = inputs[i][:x] + "O" + inputs[i][x+1:]
			break
		}
	}

	return inputs
}
