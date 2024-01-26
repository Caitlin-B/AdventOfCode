package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"reflect"
	"sort"
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

	for _, hand := range hands {
		GetHandType(hand)
	}

	// sort the hands
	sort.Slice(hands, func(i, j int) bool { return Sort(hands[i], hands[j]) })

	// sum the bet * rank
	sum := 0
	for i, hand := range hands {
		fmt.Println(hand.set, i)
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
	v := make([]int, 0, len(tally))
	for _, t := range tally {
		v = append(v, t)
	}
	// sort so uniform order
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })

	// 5 of a kind
	if reflect.DeepEqual(v, []int{5}) {
		return 7
	}
	// 4 of a kind
	if reflect.DeepEqual(v, []int{1, 4}) {
		return 6
	}
	// full house
	if reflect.DeepEqual(v, []int{2, 3}) {
		return 5
	}
	// 3 of a kind
	if reflect.DeepEqual(v, []int{1, 1, 3}) {
		return 4
	}
	// 2 pair
	if reflect.DeepEqual(v, []int{1, 2, 2}) {
		return 3
	}
	// 1 pair
	if reflect.DeepEqual(v, []int{1, 1, 1, 2}) {
		return 2
	}

	// high card
	return 1
}

var cardOrder = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13}

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
