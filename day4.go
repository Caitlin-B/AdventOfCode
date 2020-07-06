package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var passwords []int

	for i := 134792; i <= 675810; i++ {
		adjacent := false
		ascending := true
		str := strconv.Itoa(i)
		arr := strings.Split(str, "")
		for j := 1; j < len(arr); j++  {
			if arr[j] < arr[j-1] {
				ascending = false
			}
		}

		for j := 2; j < len(arr) - 1; j ++ {
			if arr[j] == arr[j-1] && arr[j] != arr[j-2] && arr[j] != arr[j+1] {
					adjacent = true
			}
		}

		if arr[0] == arr[1] && arr[1] != arr[2] {
			adjacent = true
		}

		if arr[4] == arr[5] && arr[4] != arr[3] {
			adjacent = true
		}

		if adjacent && ascending {
			passwords = append(passwords, i)
		}
	}

	fmt.Println(len(passwords))
}