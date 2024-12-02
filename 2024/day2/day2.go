package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputs := utils.ScanInput(2, 2024)
	totSafe := 0
	for _, s := range inputs {
		report := strings.Split(s, " ")
		repFl := []float64{}
		for _, r := range report {
			f, _ := strconv.ParseFloat(r, 64)
			repFl = append(repFl, f)
		}
		fmt.Println(repFl)
		if isSafe(repFl, -1) {
			fmt.Println("safe")
			totSafe++
		} else {
			for i := range repFl {
				//fmt.Println()
				//fmt.Println("itoskip, ", i)

				if isSafe(repFl, i) {
					fmt.Println("safe if removed")
					totSafe++
					break
				}
			}
		}
	}
	fmt.Println(totSafe)
}

func isSafe(in []float64, itoSkip int) bool {
	safeIncrease := 0
	safeDecrease := 0

	lasti := len(in) - 1
	if itoSkip == lasti {
		lasti--
	}
	for i := range in[0:lasti] {
		next := i + 1
		if i == itoSkip {
			continue
		}
		if next == itoSkip {
			next++
		}
		incr := in[next] - in[i]
		if math.Abs(incr) > 3 || math.Abs(incr) == 0 {
			return false
		}

		if incr < 0 {
			safeDecrease++
		} else {
			safeIncrease++
		}
	}

	if safeIncrease > 0 && safeDecrease > 0 {
		return false
	}
	return true
}
