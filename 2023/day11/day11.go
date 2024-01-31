package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
)

const expansionMultiplier = 1000000

func main() {
	inputs := utils.ScanInput(11)

	part1(inputs)
	part2(inputs)
}

func part1(in []string) {
	in = expandUniverse(in)

	galaxyPositions := getIndexOfAllGalaxies(in)
	tot := 0
	for i, pos := range galaxyPositions {
		for _, pos2 := range galaxyPositions[i+1:] {

			tot += findShortestDistBetweenTwoGalaxies(pos, pos2)
		}
	}

	fmt.Println(tot)
}

func part2(in []string) {
	in = replaceUniverse(in)

	galaxyPositions := getIndexOfAllGalaxies(in)
	tot := 0
	for i, pos1 := range galaxyPositions {
		for _, pos2 := range galaxyPositions[i+1:] {
			zerosBetweenCol := find0BetweenTwoColumns(in, pos1[1], pos1[0], pos2[0])
			zerosBetweenRow := find0BetweenTwoRows(in, pos1[0], pos1[1], pos2[1])
			tot += (zerosBetweenCol+zerosBetweenRow)*(expansionMultiplier-1) + findShortestDistBetweenTwoGalaxies(pos1, pos2)
		}
	}

	fmt.Println(tot)
}

func find0BetweenTwoColumns(in []string, rowNum, x1, x2 int) int {
	zeros := 0

	// sort smallest to largest
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	for i := x1; i < x2; i++ { // MAYBE i <= x2??
		if string(in[rowNum][i]) == "0" {
			zeros++
		}
	}
	return zeros
}

func find0BetweenTwoRows(in []string, colNum, y1, y2 int) int {
	zeros := 0

	// sort smallest to largest
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for i := y1; i < y2; i++ { // MAYBE i <= x2??
		if string(in[i][colNum]) == "0" {
			zeros++
		}
	}

	return zeros
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

// replaceUniverse replaces all rows/columns which are expandable with 0
func replaceUniverse(in []string) []string {
	rows := replaceRows(in)
	rows = replaceColumns(rows)
	return rows
}

func replaceRows(in []string) []string {
	replaces := make([]string, 0)
	nonDot := regexp.MustCompile(`[^\.+]`)
	anyC := regexp.MustCompile(`.`)

	// replace row when row is "." only with 0
	for _, row := range in {
		if nonDot.MatchString(row) {
			replaces = append(replaces, row)
			continue
		}
		replaces = append(replaces, anyC.ReplaceAllLiteralString(row, "0"))
	}
	return replaces
}

func replaceColumns(in []string) []string {
	// make a list of all columns where "." only
	indexToReplace := make([]int, 0)
	for i := 0; i < len(in[0]); i++ {
		allDots := true
		for j := range in {
			if string(in[j][i]) != "." && string(in[j][i]) != "0" {
				allDots = false
			}
		}
		if allDots {
			indexToReplace = append(indexToReplace, i)
		}
	}
	replaces := make([]string, len(in))
	// replace with "0" in each row at indexes given by indexToExpand
	for i, s := range in {
		newS := s
		for _, k := range indexToReplace {
			if k == len(newS) {
				newS = newS[0:k] + "0"
				continue
			}
			newS = newS[0:k] + "0" + newS[k+1:]
		}
		replaces[i] = newS
	}
	return replaces
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
