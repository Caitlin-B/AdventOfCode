package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func ScanInput(day, year int) []string {
    f, err := os.Open(fmt.Sprintf("./%d/day%d/data.txt", year, day))

    if err != nil {
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    f.Close()

    return lines
}

func GetNum(in string) int {
    out, err := strconv.Atoi(in)
    if err != nil {
        fmt.Println(err)
    }
    return out
}
