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
	inputs := utils.ScanInput(8, 2023)
	a := part2(inputs)
	fmt.Println(a)
}

// Finding LCM of all routes to Z
func part2(inputs []string) int {
	instructions := inputs[0]
	allSteps := getSteps(inputs)

	currentSteps := getStepKeysThatEndInA(allSteps)
	currentMaps := getMapsFromSteps(currentSteps, allSteps)
	totalSteps := make([]int, len(currentSteps))

	for i := 0; i < len(instructions); i++ {
		s := string(instructions[i])
		allStepsAtZ := true
		for j, step := range currentSteps {
			if string(step[2]) == "Z" { // reached Z leave as is
				continue
			}
			allStepsAtZ = false
			if s == "L" {
				currentSteps[j] = currentMaps[j].L
				currentMaps[j] = allSteps[currentSteps[j]]
			}
			if s == "R" {
				currentSteps[j] = currentMaps[j].R
				currentMaps[j] = allSteps[currentSteps[j]]
			}
			totalSteps[j]++
		}

		if allStepsAtZ {
			break
		}

		// if end of instructions restart
		if i == len(instructions)-1 {
			i = -1
		}
	}

	// find LCM of total steps

	return LCM(totalSteps)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	if len(integers) == 2 {
		return result
	}

	for i := 0; i < len(integers); i++ {
		result = LCM([]int{result, integers[i]})
	}

	return result
}

func getMapsFromSteps(currentSteps []string, allSteps map[string]Step) []Step {
	maps := make([]Step, 0, len(currentSteps))
	for _, s := range currentSteps {
		maps = append(maps, allSteps[s])
	}
	return maps
}

func getStepKeysThatEndInA(in map[string]Step) []string {
	endInAStep := make([]string, 0)
	for key := range in {
		if string(key[2]) == "A" {
			endInAStep = append(endInAStep, key)
		}
	}
	return endInAStep
}

func part1(inputs []string) int {
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
