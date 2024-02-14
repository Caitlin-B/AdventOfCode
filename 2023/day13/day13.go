package main

import (
    "fmt"
    "github.com/Caitlin-B/AdventOfCode/utils"
)

func main() {
    inputs := utils.ScanInput(13)
    tot := 0
    pattern := make([]string, 0)
    for _, in := range inputs {
        if in == "" {
            //if r := getRowReflection(pattern); r != 0 {
            //    fmt.Println("row: ", r)
            //    tot += r * 100
            //    pattern = []string{}
            //    continue
            //}
            c := getRowReflection(swapRowCol(pattern))
            fmt.Println("col: ", c)
            tot += c
            pattern = []string{}
            fmt.Println("------")
            continue
        }
        pattern = append(pattern, in)
    }

    fmt.Println(tot)
}

//getRowReflection - returns 0 if no horizontal reflection
func getRowReflection(in []string) int {
    //for _, i := range in {
    //    fmt.Println(i)
    //}
    for i, r := range in {
        if i != 0 && r == in[i-1] { // if row is same as row before
            firstHalf := in[:i]
            secondHalf := in[i:]
            fmt.Println("doubleRow", i)
            fmt.Println(firstHalf)
            fmt.Println(secondHalf)
            if isRowReflection(firstHalf, secondHalf) {
                return i
            }
        }
    }
    return 0
}

func isRowReflection(a, b []string) bool {
    return false
    // reverse pattern a
    a2 := a
    for i, j := 0, len(a2)-1; i < j; i, j = i+1, j-1 {
        a2[i], a2[j] = a2[j], a2[i]
    }
    length := len(a2)
    if len(b) < length {
        length = len(b)
    }
    fmt.Println(a2)
    fmt.Println(b)
    for i := 0; i < length; i++ {
        if a2[i] != b[i] {
            return false
        }
    }
    return true
}

func swapRowCol(in []string) []string {
    swapped := make([]string, 0, len(in[0]))
    for i := 0; i < len(in[0]); i++ {
        col := ""
        for j := range in {
            col = col + string(in[j][i])
        }
        swapped = append(swapped, col)
    }
    return swapped
}

// getColumnReflection - returns 0 if no vertical reflection
func getColumnReflection(in []string) int {
out:
    for i := 1; i < len(in[0]); i++ {
        for j := range in {
            if in[j][i] != in[j][i-1] {
                continue out
            }
        }
        // this point only if column equals previous column

    }

    return 0
}
