package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strings"
)

func main() {
	inputs := utils.ScanInput(10)

	initiali, initialj := 0, 0
	// get starting coords
	for k, in := range inputs {
		if strings.Contains(in, "S") {
			initialj = strings.Index(in, "S")
			initiali = k
		}
	}

	loop := true
	// i, j coords of clockwise loop
	// k, l coords of anticlockwise loop
	i, j := initiali, initialj
	// previous move, either up down left or right (U, D, L, R)
	prevMove := ""
	totSteps := 0
	for loop {
		switch string(inputs[i][j]) {
		case "S":
			i += 1 // this is manual as i have checked first step has pipe down
			prevMove = "D"
		case "|":
			switch prevMove {
			case "D":
				i += 1
			case "U":
				i -= 1
			default:
				// shouldnt reach this ever
				fmt.Println("bad step", i, j, prevMove, totSteps, "|")
			}
		case "-":
			switch prevMove {
			case "R":
				j += 1
			case "L":
				j -= 1
			default:
				fmt.Println("bad step", i, j, prevMove, totSteps, "-")
			}
		case "L":
			switch prevMove {
			case "L":
				i -= 1
				prevMove = "U"
			case "D":
				j += 1
				prevMove = "R"
			default:
				fmt.Println("bad step", i, j, prevMove, totSteps, "L")
			}
		case "J":
			switch prevMove {
			case "R":
				i -= 1
				prevMove = "U"
			case "D":
				j -= 1
				prevMove = "L"
			default:
				fmt.Println("bad step", i, j, prevMove, totSteps, "J")
			}
		case "7":
			switch prevMove {
			case "R":
				i += 1
				prevMove = "D"
			case "U":
				j -= 1
				prevMove = "L"
			default:
				fmt.Println("bad step", i, j, prevMove, totSteps, "7")
			}
		case "F":
			switch prevMove {
			case "U":
				j += 1
				prevMove = "R"
			case "L":
				i += 1
				prevMove = "D"
			}
		}

		if i == initiali && j == initialj {
			loop = false
		}
		totSteps++
	}
	fmt.Println(totSteps / 2)
}
