package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	pixels := readCSV()
	layers := make([][]string, len(pixels)/150)
	i := 0
	j := 0
	for i < len(pixels) {
		layers[j] = pixels[i:i+150]
		i += 150
		j ++
	}

	lowestZeros := []int{100,0}

	for i, layer := range layers {
		zeros := 0
		for _, pixel := range layer {
			if pixel == "0" {
				zeros ++
			}
		}

		if zeros < lowestZeros[0] {
			lowestZeros = []int{zeros, i}
		}
	}

	fmt.Println(lowestZeros)
	fewestZerosLayer := layers[lowestZeros[1]]

	ones := 0
	twos := 0

	for _, v := range fewestZerosLayer {
		if v == "1" {
			ones++
		}
		if v == "2" {
			twos++
		}
	}

	fmt.Println(ones * twos)
}

func readCSV() []string {
	file, err := os.Open("./inputs/day8.csv")
	if err != nil {
		fmt.Println("Unable to open csv") 
	}
	r:= csv.NewReader(file)
	pixelsarr, err := r.ReadAll()
	pixels := pixelsarr[0]
	if err != nil {
		fmt.Println("Unable to parse csv") 
	}

	return strings.Split(pixels[0], "")
}