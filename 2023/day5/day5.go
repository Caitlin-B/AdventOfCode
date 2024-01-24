package main

import (
	"fmt"
	"github.com/Caitlin-B/AdventOfCode/utils"
	"regexp"
)

var (
	s2s = "seed-to-soil map:"
	s2f = "soil-to-fertilizer map:"
	f2w = "fertilizer-to-water map:"
	w2l = "water-to-light map:"
	l2t = "light-to-temperature map:"
	t2h = "temperature-to-humidity map:"
	h2l = "humidity-to-location map:"
)

func main() {
	getInputs()
}

func getInputs() {
	inputs := utils.ScanInput(5)

	matchNum := regexp.MustCompile(`[0-9]+`)

	seeds := matchNum.FindAllString(inputs[0], -1)

	seedToSoil := getMap(inputs, s2s, s2f)
	soilToFertilizer := getMap(inputs, s2f, f2w)
	fertilizerToWater := getMap(inputs, f2w, w2l)
	waterToLight := getMap(inputs, w2l, l2t)
	lightToTemp := getMap(inputs, l2t, t2h)
	tempToHumid := getMap(inputs, t2h, h2l)
	humidToLoc := getMap(inputs, h2l, "")

	smallestLoc := 1000000000000000
	//seedNoOfSmallest := 0

	for _, seed := range seeds {
		seedNum := utils.GetNum(seed)
		soil := mapNumber(seedToSoil, seedNum)
		fertilizer := mapNumber(soilToFertilizer, soil)
		water := mapNumber(fertilizerToWater, fertilizer)
		light := mapNumber(waterToLight, water)
		temperature := mapNumber(lightToTemp, light)
		humidity := mapNumber(tempToHumid, temperature)
		location := mapNumber(humidToLoc, humidity)
		if location < smallestLoc {
			smallestLoc = location
			//seedNoOfSmallest = seedNum
		}
		//fmt.Println(soil)
		//fmt.Println(fertilizer)
		//fmt.Println(water)
		//fmt.Println(light)
		//fmt.Println(temperature)
		//fmt.Println(humidity)
		//fmt.Println(location)
		//fmt.Println("---------")
	}
	fmt.Println(smallestLoc)

}

// map 49 53 8, 53 should give 49 but giving 52

// maps input using mps
func mapNumber(mps []string, in int) int {
	num := 0
	for _, mp := range mps {
		m := parseMap(mp)
		if m[1] <= in && in < m[1]+m[2] {
			num = m[0] + in - m[1]

		}
	}

	if num == 0 {
		num = in
	}
	return num
}

// gets specified maps in the format ["destStart sourceStart range"]
func getMap(in []string, mapName, nextMap string) []string {
	start, end := 0, 0
	for i, v := range in {
		if v == mapName {
			start = i + 1
		}
		if v == nextMap {
			end = i - 1
		}
	}

	if nextMap == "" {
		end = len(in)
	}

	return in[start:end]
}

// parse map from "destStart sourceStart range" to [destStart, sourceStart, range]
func parseMap(str string) []int {
	matchNum := regexp.MustCompile(`[0-9]+`)

	nums := matchNum.FindAllString(str, -1)

	ints := make([]int, 3)

	for i := range nums {
		ints[i] = utils.GetNum(nums[i])
	}
	return ints
}
