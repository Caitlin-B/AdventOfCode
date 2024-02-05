package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func part2(inputs []string) int {
	sum := 0
	for _, in := range inputs {
		springs, groups := parseIn2(in)
		sum += calc2(springs, groups)
	}
	return sum
}

func parseIn2(in string) (string, []int) {
	split := strings.Split(in, " ")
	record := split[0]
	record = strings.Repeat(record+"?", 5)
	record = record[:len(record)-1]
	groupStr := strings.Repeat(split[1]+",", 5)
	groupStr = groupStr[:len(groupStr)-1]
	groupStrSpl := strings.Split(groupStr, ",")
	groupInt := make([]int, 0, len(groupStrSpl))
	for _, g := range groupStrSpl {
		i, err := strconv.Atoi(g)
		if err != nil {
			fmt.Println(err.Error())
		}
		groupInt = append(groupInt, i)
	}

	return record, groupInt
}

func getKey(springs string, groups []int) string {
	return springs + fmt.Sprintf("%v", groups)
}

func calc2(springs string, groups []int) int {
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
		key := getKey(springs[1:], groups)
		c, ok := cache[key]
		if !ok {
			c = calc2(springs[1:], groups)
			cache[key] = c
		}
		sum = c

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

	key := getKey(springs[groups[0]:], groups[1:])
	c, ok := cache[key]
	if !ok {
		c = calc2(springs[groups[0]:], groups[1:])
		cache[key] = c
	}

	return sum + c
}
