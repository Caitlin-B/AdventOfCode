package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strings"
)

const lenFirstSect = 5

func main() {
	inputs := utils.ScanInput(4)

	sum := 0
	for _, i := range inputs {
		winners, got := getNums(i)
		val := cardVal(winners, got)

		sum += val
		fmt.Println("cardval: ", val)
		fmt.Println("-----------")
	}
	fmt.Println(sum)
}

func getNums(in string) (winners []string, got []string) {
	arr := strings.Split(in, " ")
	fmt.Println(arr, len(arr))
	arr = removeSpace(arr)
	fmt.Println(arr, len(arr))

	pipe := indexOfStr(arr, "|")
	winners = arr[2:pipe]
	got = arr[pipe+1:]
	fmt.Println(winners, got)
	fmt.Println(len(winners), len(got))

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
