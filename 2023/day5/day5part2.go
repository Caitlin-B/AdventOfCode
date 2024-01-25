package main

import (
    "fmt"
    "github.com/Caitlin-B/AdventOfCode/utils"
    "regexp"
)

func part2() {
    inputs := utils.ScanInput(5)
    matchNum := regexp.MustCompile(`[0-9]+`)

    seeds := matchNum.FindAllString(inputs[0], -1)

    r := getRanges(seeds)

    cont := true
    i := 0
    for cont {
        if i%100 == 0 {
            fmt.Println(i)
        }
        toSeed := locationToSeed(inputs, i)
        if isInRanges(toSeed, r) {
            cont = false
        }
        i++
    }
    fmt.Println("seed within range: ", i)
}

func getRanges(seeds []string) (ranges []Range) {
    for i := 0; i < len(seeds); i += 2 {
        start := utils.GetNum(seeds[i])
        end := start + utils.GetNum(seeds[i+1])
        ranges = append(ranges, Range{start: start, end: end})
    }
    return
}

// checks whether int is within ranges
func isInRanges(i int, ranges []Range) bool {
    for _, r := range ranges {
        if r.start <= i && i < r.end {
            return true
        }
    }

    return false
}

func locationToSeed(inputs []string, i int) int {
    seedToSoil := getMap(inputs, s2s, s2f)
    soilToFertilizer := getMap(inputs, s2f, f2w)
    fertilizerToWater := getMap(inputs, f2w, w2l)
    waterToLight := getMap(inputs, w2l, l2t)
    lightToTemp := getMap(inputs, l2t, t2h)
    tempToHumid := getMap(inputs, t2h, h2l)
    humidToLoc := getMap(inputs, h2l, "")

    humidity := mapNumberReverse(humidToLoc, i)
    temperature := mapNumberReverse(tempToHumid, humidity)
    light := mapNumberReverse(lightToTemp, temperature)
    water := mapNumberReverse(waterToLight, light)
    fertilizer := mapNumberReverse(fertilizerToWater, water)
    soil := mapNumberReverse(soilToFertilizer, fertilizer)
    seed := mapNumberReverse(seedToSoil, soil)

    return seed
}

// maps in reverse
func mapNumberReverse(mps []string, in int) int {
    num := 0
    for _, mp := range mps {
        m := parseMap(mp)
        if m[0] <= in && in < m[0]+m[2] {
            num = m[1] + in - m[0]

        }
    }

    if num == 0 {
        num = in
    }
    return num
}
