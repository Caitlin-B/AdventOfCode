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
		if isSafe(repFl) {
			totSafe++
		}
	}
	fmt.Println(totSafe)
}

func isSafe(in []float64) bool {
	safeIncrease := 0
	safeDecrease := 0

	for i := range in[0 : len(in)-1] {
		incr := in[i+1] - in[i]
		if math.Abs(incr) > 3 || math.Abs(incr) == 0 {
			fmt.Println("big increase", in, in[i], " to ", in[i+1])
			return false
		}
		if incr < 0 {
			safeDecrease++
		} else {
			safeIncrease++
		}
	}

	if safeIncrease > 0 && safeDecrease > 0 {
		fmt.Println("not all increasing", in)
		return false
	}
	return true
}
