package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"slices"
	"strings"
)

const rulesEnd = 1176

func main() {
	inputs := utils.ScanInput(5, 2024)

	rules := make(map[string][]string, rulesEnd)
	for _, in := range inputs[:rulesEnd] {
		n1 := in[:2]
		n2 := in[3:]
		if ns, ok := rules[n1]; !ok {
			rules[n1] = []string{n2}
		} else {
			rules[n1] = append(ns, n2)
		}
	}

	tot := 0
	for _, in := range inputs[rulesEnd+1:] {
		updates := strings.Split(in, ",")
		//tot += isCorrect(updates, rules)
		if isCorrect(updates, rules) == 0 {
			tot += orderPage(updates, rules)
		}
	}

	fmt.Println(tot)
}

func isCorrect(update []string, rules map[string][]string) int {
	for i, page := range update {
		if i == 0 {
			continue
		}
		mustAfter := rules[page]

		for _, must := range mustAfter {
			if slices.Index(update[:i], must) != -1 {
				return 0
			}
		}
	}
	middle := update[len(update)/2]
	return utils.GetNum(middle)
}

func orderPage(update []string, rules map[string][]string) int {
	slices.SortFunc(update, func(a, b string) int {
		return slices.Index(rules[a], b)
	})
	middle := update[len(update)/2]
	return utils.GetNum(middle)
}
