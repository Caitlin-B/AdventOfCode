package main

import (
    "fmt"
    "github.com/Caitlin-B/AdventOfCode/utils"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    inputs := utils.ScanInput(2)

    tot := 0
    for _, game := range inputs {
        red := colourPass(game, "red", 12)
        green := colourPass(game, "green", 13)
        blue := colourPass(game, "blue", 14)

        tot += blue * green * red
    }

    fmt.Println(tot)
}

func colourPass(game, colour string, max int) int {
    reg := regexp.MustCompile(`[0-9]+ ` + colour)

    counts := reg.FindAllString(game, -1)
    min := 0
    for _, count := range counts {
        c := strings.Trim(count, " "+colour)
        t, err := strconv.Atoi(c)
        if err != nil {
            fmt.Println(err.Error())
            return 0
        }
        if t > min {
            min = t
        }
    }

    return min
}
