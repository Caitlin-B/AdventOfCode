package main

import (
    "fmt"
    "github.com/Caitlin-B/AdventOfCode/utils"
    "regexp"
    "strconv"
)

func main() {
    part1()
}

func part1() {
    inputs := utils.ScanInput(3)

    matchNum := regexp.MustCompile(`[0-9]+`)
    matchSymbol := regexp.MustCompile(`[^0-9.]`)

    tot := 0
    for rowIndex, row := range inputs {
        // find indexes of numbers
        numbers := matchNum.FindAllString(row, -1)
        indexesOfNumbers := matchNum.FindAllStringIndex(row, -1)

        fmt.Println(indexesOfNumbers)
        fmt.Println(numbers)

        // for all indexes
        for i, ind := range indexesOfNumbers {
            nextToSymbol := false

            firstIOfNumber := ind[0]
            lastIOfNumber := ind[1]

            // for searching adjacent rows
            startSearchIndex := firstIOfNumber
            endSearchIndex := lastIOfNumber

            // find character left of match
            if firstIOfNumber != 0 {
                startSearchIndex = firstIOfNumber - 1
                charOnLeft := row[startSearchIndex]
                nextToSymbol = nextToSymbol || matchSymbol.MatchString(string(charOnLeft))
                //fmt.Println("regex: ", matchSymbol.MatchString(string(charOnLeft)))
                //fmt.Println("next to symbol: ", nextToSymbol)
            }
            // find character right of match
            if lastIOfNumber != len(row) {
                //endSearchIndex = lastIOfNumber + 1
                charOnRight := row[endSearchIndex]
                //fmt.Println("charonRight: ", string(charOnRight))

                nextToSymbol = nextToSymbol || matchSymbol.MatchString(string(charOnRight))
                //fmt.Println("regex: ")
                //fmt.Println(matchSymbol.MatchString(string(charOnRight)))
                //fmt.Println("next to symbol:")
                //fmt.Println(nextToSymbol)
                // for next step
                endSearchIndex++
            }

            //find above chars (diag incl)
            if rowIndex != 0 {
                chars := inputs[rowIndex-1][startSearchIndex:endSearchIndex]
                nextToSymbol = nextToSymbol || matchSymbol.MatchString(chars)
                fmt.Println("chars: " + chars)
                fmt.Println("regex: ")
                fmt.Println(matchSymbol.MatchString(string(chars)))
                fmt.Println("next to symbol:")
                fmt.Println(nextToSymbol)
            }
            // find bellow chars (diag incl)
            if rowIndex != len(inputs)-1 { // 139 for final
                chars := inputs[rowIndex+1][startSearchIndex:endSearchIndex]
                nextToSymbol = nextToSymbol || matchSymbol.MatchString(chars)
                fmt.Println("regex: ")
                fmt.Println(matchSymbol.MatchString(string(chars)))
                fmt.Println("next to symbol:")
                fmt.Println(nextToSymbol)
            }

            if nextToSymbol == true {
                num, err := strconv.Atoi(numbers[i])
                if err != nil {
                    fmt.Println(err)
                }
                tot += num
            }

        }
        fmt.Println(tot)

        fmt.Println("---------")
    }
}
