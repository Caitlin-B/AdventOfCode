package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//read csv of opcode
	file, err := os.Open("./inputs/day3.csv")
	if err != nil {
		fmt.Println("Unable to open csv") 
	}
	r:= csv.NewReader(file)
	patharr, err := r.ReadAll()
	path1 := patharr[0]
	path2 := patharr[1]
	if err != nil {
		fmt.Println("Unable to parse csv") 
	}

	//path1 = []string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"}
	//path2 = []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"}
	//calculate every step a path takes 
	//calcPoints gives an array of each point the path crosses
	points1 := calcPoints(path1)
	points2 := calcPoints(path2)


	//find an array of each intersecting point
	var intersections [][]int
	
	for i := 0; i < len(points1); i++ {
		for j := 0; j < len(points2); j++ {
			if points1[i][0] == points2[j][0] && points1[i][1] == points2[j][1] {
				intersections = append(intersections, points1[i])
			}
		}
	}

	fmt.Println("INTERSECTIONS", intersections[1:])

	//find manhattan distance of each intersection
	var manhattans []int

	for _, v := range intersections {
		distance := v[0] + v[1]
		manhattans = append(manhattans, distance)
	}

	//fmt.Println("MANHATTANS", manhattans)

	//calc steps taken to intersections for each path
	fmt.Println(".")
	//fmt.Println(points1[1:])
	//fmt.Println(points2[1:])
	//fmt.Println(".")
	steps1 := calcSteps(points1[1:], intersections[1:])
	steps2 := calcSteps(points2[1:], intersections[1:])

	fmt.Println("STEPS 1", steps1) 
	fmt.Println("STEPS 2", steps2)

	var totalSteps []int

	for i := range steps1 {
		sum:= steps1[i] + steps2[i] + 2
		totalSteps = append(totalSteps, sum)
	}
	fmt.Println(".")
	fmt.Println("TOTAL STEPS", totalSteps)


}

func calcSteps(points [][]int, intersections[][]int) []int{
	//calc how many steps it takes to get to each intersection
	var steps []int
	for _, intersect := range intersections {
		for i, point := range points {
			if point[0] == intersect[0] && point[1] == intersect[1] {
				steps = append(steps, i)
			}
		}
	}

	return steps
}

func calcPoints(path []string) [][]int {
	var points [][]int
	points = append(points, []int{0, 0})
	for _, v := range path {
		direction := string(v[0])
		steps, err := strconv.Atoi(v[1:])
		if err != nil {fmt.Println("fuck it")}

		//take last position
		prevPos := points[len(points) - 1]

		//move steps according to direction, add to points
		if direction == "L" {
			for i := 1; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[0] = newPos[0] - i
				points = append(points, newPos)
			}
		} else if direction == "R" {
			for i := 1; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[0] = newPos[0] + i
				points = append(points, newPos)
			}
		} else if direction == "U" {
			for i := 1; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[1] = newPos[1] + i
				points = append(points, newPos)
			}
		} else if direction == "D" {
			for i := 1; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[1] = newPos[1] - i
				points = append(points, newPos)
				
			}
		}
	}

	return points
}