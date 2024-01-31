package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
)

func main() {
	inputs := utils.ScanInput(11)
	inputs = expandUniverse(inputs)
	//numbered, totalGal := assignNumbersToGalaxies(inputs)

	galaxyPositions := getIndexOfAllGalaxies(inputs)
	tot := 0
	for i, pos := range galaxyPositions {
		for _, pos2 := range galaxyPositions[i+1:] {

			tot += findShortestDistBetweenTwoGalaxies(pos, pos2)
		}
	}

	fmt.Println(tot)
}

func findShortestDistBetweenTwoGalaxies(gal1, gal2 []int) int {
	xdiff := gal1[0] - gal2[0]
	if xdiff < 0 {
		xdiff = -xdiff
	}

	ydiff := gal1[1] - gal2[1]
	if ydiff < 0 {
		ydiff = -ydiff
	}
	return xdiff + ydiff
}

// get x and y coord of each galaxy
func getIndexOfAllGalaxies(in []string) [][]int {
	gPos := make([][]int, 0)
	hash := regexp.MustCompile(`[#]`)

	for i, r := range in {
		s := hash.FindAllStringIndex(r, -1)

		for _, j := range s {
			gPos = append(gPos, []int{j[0], i})
		}
	}

	return gPos
}

func expandUniverse(in []string) []string {
	rows := expandRows(in)
	rows = expandColumns(rows)
	return rows
}

func expandRows(in []string) []string {
	expanded := make([]string, 0)
	nonDot := regexp.MustCompile(`[^\.+]`)

	// expand extra row when row is "." only
	for _, row := range in {
		if nonDot.MatchString(row) {
			expanded = append(expanded, row)
			continue
		}
		expanded = append(expanded, row)
		expanded = append(expanded, row)
	}
	return expanded
}

func expandColumns(in []string) []string {
	// make a list of all columns where "." only
	indexToExpand := make([]int, 0)
	for i := 0; i < len(in[0]); i++ {
		allDots := true
		for j := range in {
			if string(in[j][i]) != "." {
				allDots = false
			}
		}
		if allDots {
			indexToExpand = append(indexToExpand, i)
		}
	}

	expanded := make([]string, len(in))
	// insert "." in each row at indexes given by indexToExpand
	for i, s := range in {
		newS := s
		for j, k := range indexToExpand {
			newS = newS[0:k+j] + "." + newS[k+j:]
		}
		expanded[i] = newS
	}
	return expanded
}
