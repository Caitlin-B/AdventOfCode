package main

import "fmt"

type Race struct {
	time     int
	distance int
}

var races = []Race{{48938466, 261119210191063}}

func main() {
	tot := part1(races)

	fmt.Println("final: ", tot)
}

func part1(races []Race) int {
	tot := 1
	for _, race := range races {
		winners := 0
		for i := 0; i < race.time; i++ {
			totalDist := (race.time - i) * i
			if totalDist > race.distance {
				winners++
			}
		}
		tot *= winners
	}
	return tot
}
