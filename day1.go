package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//read csv of module mass
	file, err := os.Open("./inputs/day1.csv")
	if err != nil {
		fmt.Println("Unable to open csv") //means something went wrong
	}
	r:= csv.NewReader(file)
	modules, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse csv") //means something went wrong
	}
	
	//calculate fuel requirements
	
	for _, mass := range modules {	
		if s, err := strconv.ParseFloat(mass[0], 32); err == nil {
			calculateFuel(s)
		}
	}

	fmt.Println(sum)
}

var sum float64 = 0

func calculateFuel(mass float64) float64 {
	fuelNeeded := math.Floor(mass / 3) - 2

	if fuelNeeded >= 0 {
		sum += fuelNeeded
		calculateFuel(fuelNeeded)
	}

	return sum
}