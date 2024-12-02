package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strings"
)

func main() {
	//part1()
	part2()
}

// slice of card matches
// slice of card copies

func part2() {
	inputs := utils.ScanInput(4, 2023)

	matchesOnCards := make([]int, len(inputs))
	cardCopies := make([]int, len(inputs))
	for i, card := range inputs {
		matchesOnCards[i] = getNumberOfMatches(card)
		cardCopies[i] = 1
	}

	for i, matches := range matchesOnCards {
		for j := 0; j < cardCopies[i]; j++ {
			for k := i + 1; k < i+1+matches; k++ {
				cardCopies[k] = cardCopies[k] + 1
			}
		}
	}

	tot := 0
	for _, n := range cardCopies {
		tot += n
	}

	fmt.Println(tot)
}

func part1() {
	inputs := utils.ScanInput(4, 2023)

	sum := 0
	for _, i := range inputs {
		winners, got := getNums(i)
		val := cardVal(winners, got)

		sum += val
	}
	fmt.Println(sum)
}

func getNumberOfMatches(in string) int {
	winners, got := getNums(in)
	return cardMatches(winners, got)
}

func getNums(in string) (winners []string, got []string) {
	arr := strings.Split(in, " ")
	//fmt.Println(arr, len(arr))
	arr = removeSpace(arr)
	//fmt.Println(arr, len(arr))

	pipe := indexOfStr(arr, "|")
	winners = arr[2:pipe]
	got = arr[pipe+1:]
	//fmt.Println(winners, got)
	//fmt.Println(len(winners), len(got))

	return
}

func cardVal(w, g []string) int {
	count := 0
	for _, n := range g {
		if indexOfStr(w, n) != -1 {
			if count == 0 {
				count = 1
			} else {
				count = count * 2
			}
		}
	}
	return count
}
func cardMatches(w, g []string) int {
	count := 0
	for _, n := range g {
		if indexOfStr(w, n) != -1 {
			count++
		}
	}
	return count
}

func indexOfStr(slice []string, str string) int {
	for i, s := range slice {
		if s == str {
			return i
		}
	}
	return -1
}

func removeSpace(slice []string) []string {
	newS := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != "" {
			newS = append(newS, s)
		}
	}
	return newS
}
