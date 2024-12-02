package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	set string
	bid int
}

func main() {
	fmt.Println(part1())
}

func part1() int {
	inputs := utils.ScanInput(7, 2023)
	hands := inputToHands(inputs)

	for _, hand := range hands {
		GetHandType(hand)
	}

	// sort the hands
	sort.Slice(hands, func(i, j int) bool { return Sort(hands[i], hands[j]) })

	// sum the bet * rank
	sum := 0
	for i, hand := range hands {
		//fmt.Println(hand.set, i)
		sum += (i + 1) * hand.bid
	}

	return sum
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

func Sort(first, second Hand) bool {
	type1, type2 := GetHandType(first), GetHandType(second)

	if type1 == type2 {
		return SortHighCard(first, second)
	}

	return type1 < type2
}

// GetHandType gets the value of the hand type as follows:
func GetHandType(hand Hand) int {
	tally := map[string]int{}
	for _, c := range hand.set {
		if _, ok := tally[string(c)]; ok {
			tally[string(c)]++
			continue
		}
		tally[string(c)] = 1
	}

	// slice of card counts
	c := make([]int, 0, len(tally))
	for _, t := range tally {
		c = append(c, t)
	}

	// sort so uniform order
	sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })

	// DAY2
	// get number of any jokers
	jokers := tally["J"]
	// add to the total of the biggest card count (that isnt a joker) and remove joker from count
	jokerRemoved := false // whether the joker value has been removed from the slice
	jokerAdded := false   // whether the joker value has been added to another value
	newc := make([]int, 0, len(tally))
	for _, v := range c {
		newv := 0
		if v == jokers && jokerRemoved && jokerAdded {
			newv = v
		}
		if v == jokers && jokerRemoved && !jokerAdded {
			newv = v + jokers
			jokerAdded = true
		}
		if v == jokers && !jokerRemoved {
			// dont append
			jokerRemoved = true
			continue
		}
		if v != jokers && jokerAdded {
			newv = v
		}
		if v != jokers && !jokerAdded {
			newv = v + jokers
			jokerAdded = true
		}
		newc = append(newc, newv)
	}
	// \DAY2

	// handle edge case of all cards joker, which is 5 of a kind
	if len(newc) == 0 {
		return 7
	}
	// 5 of a kind
	if newc[0] == 5 {
		return 7
	}
	// 4 of a kind
	if newc[0] == 4 {
		return 6
	}
	// full house
	if len(newc) > 1 && newc[0] == 3 && newc[1] == 2 {
		return 5
	}
	// 3 of a kind
	if newc[0] == 3 {
		return 4
	}
	// 2 pair
	if len(newc) > 1 && newc[0] == 2 && newc[1] == 2 {
		return 3
	}
	// 1 pair
	if newc[0] == 2 {
		return 2
	}

	return 1
}

var cardOrder = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"T": 4,
	"9": 5,
	"8": 6,
	"7": 7,
	"6": 8,
	"5": 9,
	"4": 10,
	"3": 11,
	"2": 12,
	"J": 13}

// SortHighCard returns true of first hand is higher than second
func SortHighCard(first, second Hand) bool {
	for i := 0; i < 5; i++ {
		if first.set[i] == second.set[i] {
			continue
		}
		return cardOrder[string(first.set[i])] > cardOrder[string(second.set[i])]
	}
	// should not reach this point
	return false
}
