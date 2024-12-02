package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	inputs := utils.ScanInput(1, 2024)
	part2(inputs)

}

func part2(inputs []string) {
	list1 := []int{}
	list2 := map[int]int{}
	for _, s := range inputs {
		r, _ := regexp.Compile("\\d+")
		mat := r.FindAllString(s, -1)
		i1, _ := strconv.Atoi(mat[0])
		list1 = append(list1, i1)

		i2, _ := strconv.Atoi(mat[1])
		list2[i2]++
	}
	tot := 0
	for _, in := range list1 {
		tot += list2[in] * in
	}

	fmt.Println(tot)
}

func part1(inputs []string) {
	list1 := []int{}
	list2 := []int{}
	for _, s := range inputs {
		r, _ := regexp.Compile("\\d+")
		mat := r.FindAllString(s, -1)
		i1, _ := strconv.Atoi(mat[0])
		list1 = append(list1, i1)
		i2, _ := strconv.Atoi(mat[1])
		list2 = append(list2, i2)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	tot := 0
	for i := range list1 {
		tot += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(tot)
}
