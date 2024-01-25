package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"strconv"
	"strings"
)

type Hand struct {
	set string
	bid int
}

func main() {
	fmt.Println(day1())
}

func day1() int {
	inputs := utils.ScanInput(7)
	hands := inputToHands(inputs)

	_ = hands
	return 0
}

func inputToHands(inputs []string) []Hand {
	hands := make([]Hand, 0, len(inputs))
	for _, in := range inputs {
		ins := strings.Split(in, " ")
		bid, err := strconv.Atoi(ins[1])
		if err != nil {
			fmt.Println(err)
		}
		hands = append(hands, Hand{set: ins[0], bid: bid})
	}
	return hands
}

// GetHandType gets the value of the hand type as follows:
// high card = 1
// one pair = 2
// two pair = 3
// three of a kind = 4
// full house = 5
// four of a kind = 6
// five of a kind = 7
func GetHandType(hand Hand) int {

}
