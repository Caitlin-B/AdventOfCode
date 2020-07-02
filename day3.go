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

	fmt.Println(path1[0], "IGNORE")
	fmt.Println(path2[0], "IGNORE")

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

	fmt.Println(intersections)

	//find manhattan distance of each intersection
	var manhattans []int

	for _, v := range intersections {
		distance := v[0] + v[1]
		manhattans = append(manhattans, distance)
	}

	fmt.Println(manhattans)

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
			for i := 0; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[0] = newPos[0] - i
				points = append(points, newPos)
			}
		} else if direction == "R" {
			for i := 0; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[0] = newPos[0] + i
				points = append(points, newPos)
			}
		} else if direction == "U" {
			for i := 0; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[1] = newPos[1] + i
				points = append(points, newPos)
			}
		} else if direction == "D" {
			for i := 0; i <= steps; i++ {
				newPos := make([]int, 2)
				copy(newPos, prevPos)
				newPos[1] = newPos[1] - i
				points = append(points, newPos)
				
			}
		}
	}

	return points
}