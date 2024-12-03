package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
	"strconv"
)

func main() {
	inputs := utils.ScanInput(3, 2024)
	regMul, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	regD, _ := regexp.Compile(`\d{1,3}`)
	muls := regMul.FindAllString(inputs[0], -1)
	tot := 0
	for _, m := range muls {
		is := regD.FindAllString(m, -1)
		a, _ := strconv.Atoi(is[0])
		b, _ := strconv.Atoi(is[1])
		//fmt.Println(m, a, b)

		tot += a * b
	}
	fmt.Println(tot)
}
