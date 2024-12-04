package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
	"strconv"
)

func main() {
	inputs := utils.ScanInput(3, 2024)
	regMul, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	regD, _ := regexp.Compile(`\d{1,3}`)
	muls := regMul.FindAllString(inputs[0], -1)
	tot := 0
	do := true
	for _, m := range muls {
		if m == "do()" {
			do = true
			continue
		}
		if m == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		is := regD.FindAllString(m, -1)
		a, _ := strconv.Atoi(is[0])
		b, _ := strconv.Atoi(is[1])
		//fmt.Println(m, a, b)

		tot += a * b
	}
	fmt.Println(tot)

}
