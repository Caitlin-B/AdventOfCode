package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
	"strconv"
)

func main() {
	inputs := utils.ScanInput(1, 2023)
	reg := regexp.MustCompile(`[^0-9]+`)

	total := 0
	for _, in := range inputs {
		wordsToNums := lettersToNumbers(in)
		nums := reg.ReplaceAllLiteralString(wordsToNums, "")

		lastIndex := 0

		if len(nums) != 1 {
			lastIndex = len(nums) - 1
		}

		num, _ := strconv.Atoi(string(nums[0]) + string(nums[lastIndex]))
		fmt.Println(num)
		fmt.Println("-------")
		total += num
	}

	fmt.Println(total)
}

func lettersToNumbers(str string) string {
	fmt.Println(str)

	for i := 0; i < len(str)-4; i++ {
		end := i + 5
		if i > len(str)-5 {
			end = len(str) - i
		}
		replace := replaceAllNumWords(str[i:end])
		str = str[:i] + replace + str[end:]
	}

	fmt.Println(str)
	return str
}
func replaceAllNumWords(str string) string {
	str = letterToNumber(str, `one`, "1ne")
	str = letterToNumber(str, `two`, "2wo")
	str = letterToNumber(str, `three`, "3hree")
	str = letterToNumber(str, `four`, "4our")
	str = letterToNumber(str, `five`, "5ive")
	str = letterToNumber(str, `six`, "6ix")
	str = letterToNumber(str, `seven`, "7even")
	str = letterToNumber(str, `eight`, "8ight")
	str = letterToNumber(str, `nine`, "9ine")
	return str
}
func letterToNumber(str, numStr, numnum string) string {
	reg := regexp.MustCompile(numStr)
	return reg.ReplaceAllLiteralString(str, numnum)

}
