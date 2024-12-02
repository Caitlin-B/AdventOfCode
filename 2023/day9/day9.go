package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
	"strconv"
)

func main() {
	inputs := utils.ScanInput(9, 2023)
	tot := 0
	for _, in := range inputs {
		inInt := inputToInt(in)
		// for part2 reverse slice
		for i, j := 0, len(inInt)-1; i < j; i, j = i+1, j-1 {
			inInt[i], inInt[j] = inInt[j], inInt[i]
		}

		_, h := getDifferenceLast(inInt, inInt[len(inInt)-1])
		tot += h
	}
	fmt.Println(tot)

}

func inputToInt(in string) []int {
	matchNum := regexp.MustCompile(`-*[0-9]+`)
	numStrings := matchNum.FindAllString(in, -1)
	numInts := make([]int, 0, len(numStrings))

	for _, n := range numStrings {
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err.Error())
		}
		numInts = append(numInts, i)
	}

	return numInts
}

func getDifferenceLast(in []int, totDiff int) ([]int, int) {
	diffs := make([]int, 0, len(in)-1)
	allZero := true
	for i := range in {
		if in[i] != 0 {
			allZero = false
		}
		if i == 0 { //skip first
			continue
		}
		diff := in[i] - in[i-1]
		//if diff < 0 {
		//	diff = -diff
		//}
		diffs = append(diffs, diff)
	}

	if allZero {
		return nil, totDiff
	}
	//totdiff add total of last in slice
	return getDifferenceLast(diffs, totDiff+diffs[len(diffs)-1])
}
