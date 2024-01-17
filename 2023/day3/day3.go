package main

import (
    "fmt"
    "github.com/Caitlin-B/AdventOfCode/utils"
    "regexp"
    "strconv"
)

func main() {
    //part1()
    part2()
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
            }
            // find bellow chars (diag incl)
            if rowIndex != len(inputs)-1 { // 139 for final
                chars := inputs[rowIndex+1][startSearchIndex:endSearchIndex]
                nextToSymbol = nextToSymbol || matchSymbol.MatchString(chars)
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

func part2() {
    inputs := utils.ScanInput(3)

    matchStar := regexp.MustCompile(`\*`)
    matchNum := regexp.MustCompile(`[0-9]+`)

    tot := 0
    for rowIndex, row := range inputs {
        // find indexes of star
        indexesOfStars := matchStar.FindAllStringIndex(row, -1)
        // for all indexes
        for _, ind := range indexesOfStars {
            adjNumbers := make([]string, 0, 6)

            indexOfStar := ind[0]

            startSearchIndex := 0
            if indexOfStar > 3 {
                startSearchIndex = indexOfStar - 3
            }
            endSearchIndex := ind[1] + 3
            if endSearchIndex >= len(row) {
                endSearchIndex = len(row) - 1
            }

            // find number left of star
            if indexOfStar != 0 && matchNum.MatchString(string(row[indexOfStar-1])) {
                numLeft := matchNum.FindString(row[startSearchIndex:indexOfStar])
                adjNumbers = append(adjNumbers, numLeft)
                fmt.Println("LEFTNUM", numLeft)
            }
            // find number right of star
            if indexOfStar != len(row) && matchNum.MatchString(string(row[indexOfStar+1])) {
                numRight := matchNum.FindString(row[indexOfStar:endSearchIndex])
                adjNumbers = append(adjNumbers, numRight)
                fmt.Println("RIGHTNUM", numRight)
            }

            // find number above star
            if rowIndex != 0 {
                numsAbove := getNumsInAdgRow(inputs[rowIndex-1][startSearchIndex:endSearchIndex])
                adjNumbers = append(adjNumbers, numsAbove...)
            }

            // find number below star
            if rowIndex != len(inputs)-1 {
                numsBelow := getNumsInAdgRow(inputs[rowIndex+1][startSearchIndex:endSearchIndex])
                adjNumbers = append(adjNumbers, numsBelow...)
            }

            fmt.Println("ADJNUMS", adjNumbers)

            if len(adjNumbers) == 2 {
                num1, err := strconv.Atoi(adjNumbers[0])
                if err != nil {
                    fmt.Println(err.Error())
                }
                num2, err := strconv.Atoi(adjNumbers[1])
                if err != nil {
                    fmt.Println(err.Error())
                }

                tot += num1 * num2
            }

        }
    }
    fmt.Println(tot)

    fmt.Println("---------")
}

// str should be 7 long
func getNumsInAdgRow(str string) []string {
    matchNum := regexp.MustCompile(`[0-9]+`)
    // split into numbers removing dots
    nums := matchNum.FindAllString(str, -1)
    idx := matchNum.FindAllStringIndex(str, -1)
    // no numbers
    if len(nums) == 0 {
        return []string{}
    }

    adjNums := make([]string, 0, 2)

    // check each number occupies middle three points (thus adj to *)
    for i, num := range nums {
        // if starts between 1 and 4 or finishes between 2 and 5
        startAfter1 := 1 <= idx[i][0]
        startBefore4 := idx[i][0] <= 4
        endAfter3 := 3 <= idx[i][1]
        endBefore5 := idx[i][1] <= 5
        if (startAfter1 && startBefore4) || (endAfter3 && endBefore5) {
            adjNums = append(adjNums, num)
        }
    }

    return adjNums
}
