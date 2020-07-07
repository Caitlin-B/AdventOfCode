package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	image := readCSV()
	layers := make([][]string, len(image)/150)
	i := 0
	j := 0
	for i < len(image) {
		layers[j] = image[i:i+150]
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

	pixels := make([][]string, 150)

	for i := range layers[0] {
		for _, v := range layers {
			pixels[i] = append(pixels[i], v[i])
		}
	}

	for i, v1 := range pixels {
		var newPixel string
		for _, v2 := range v1 {
			if v2 == "0" {
				newPixel = string('█')
				break
			} else if v2 == "1" {
				newPixel = "1"
				break
			}
		}
		pixels[i] = []string{newPixel}
	}

	
	formatPixels := make([]string, len(pixels))

	for i, v := range pixels {
		formatPixels[i] = v[0]
	}

	fmt.Println(formatPixels)
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