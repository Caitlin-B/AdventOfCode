package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
)

type Step struct {
	L string
	R string
}

func main() {
	inputs := utils.ScanInput(8)
	a := day1(inputs)
	fmt.Println(a)
}

func day1(inputs []string) int {
	instructions := inputs[0]
	steps := getSteps(inputs)

	currentStep := "AAA"
	currentMap := steps[currentStep]
	totalSteps := 0

	for i := 0; i < len(instructions); i++ {
		s := string(instructions[i])
		//fmt.Println(currentStep, currentMap, s)
		if currentStep == "ZZZ" {
			return totalSteps
		}
		if s == "L" {
			currentStep = currentMap.L
			currentMap = steps[currentStep]
		}
		if s == "R" {
			currentStep = currentMap.R
			currentMap = steps[currentStep]
		}
		// if end of instructions restart
		if i == len(instructions)-1 {
			i = -1
		}

		totalSteps++
	}

	return totalSteps
}

func getSteps(inputs []string) map[string]Step {
	steps := make(map[string]Step)

	for _, s := range inputs[2:] {
		key := s[0:3]
		steps[key] = Step{
			L: s[7:10],
			R: s[12:15],
		}
	}

	return steps
}
