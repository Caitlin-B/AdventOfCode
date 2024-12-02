package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strconv"
	"strings"
)

const (
	operational = "."
	damaged     = "#"
	unknown     = "?"
)

func main() {
	inputs := utils.ScanInput(12, 2023)

	//sum := part1(inputs)
	sum := part2(inputs)
	fmt.Println(sum)
}

func part1(inputs []string) int {
	sum := 0
	for _, in := range inputs {
		record, groups := parseIn(in)
		sum += calc(record, groups)
	}
	return sum
}

func parseIn(in string) (string, []int) {
	split := strings.Split(in, " ")
	record := split[0]
	groupStr := strings.Split(split[1], ",")
	groupInt := make([]int, 0, len(groupStr))
	for _, g := range groupStr {
		i, err := strconv.Atoi(g)
		if err != nil {
			fmt.Println(err.Error())
		}
		groupInt = append(groupInt, i)
	}

	return record, groupInt
}

func calc(springs string, groups []int) int {
	springs = strings.Trim(springs, operational)

	// if no springs left, confirm no groups left, otherwise not a valid arrangement
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	// if no groups left, confirm no damaged springs left, otherwise not valid arrangement
	if len(groups) == 0 {
		if strings.Contains(springs, damaged) {
			return 0
		}
		return 1
	}

	// if not enough to fulfil next group, not valid
	if len(springs) < groups[0] {
		return 0
	}

	sum := 0

	// if first char is unknown - evaluate as either operational or damaged
	if string(springs[0]) == unknown {
		// calc as operational
		sum = calc(springs[1:], groups)
		// calc as damaged
		springs = damaged + springs[1:]
	}

	// if first char damaged - must be enough to fulfil first group.
	// if it is, continue calc with remaining springs and groups
	//if string(springs[0]) == damaged { // should always be damaged at this point
	firstGroup := springs[:groups[0]]
	firstGroup = strings.Replace(firstGroup, unknown, damaged, -1)
	if strings.Contains(firstGroup, ".") { // group cannot be fulfilled
		return sum
	}

	// if next spring...
	if len(springs) > groups[0] {
		// ... is damaged group cannot be fulfilled
		if string(springs[groups[0]]) == damaged {
			return sum
		}
		// otherwise must be operational
		springs = springs[:groups[0]] + operational + springs[groups[0]+1:]
	}

	return sum + calc(springs[groups[0]:], groups[1:])
}
