package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func ScanInput(day int) []string {
    file, err := os.Open(fmt.Sprintf("./2023/day%d/data.txt", day))

    if err != nil {
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    file.Close()

    return lines
}

func GetNum(in string) int {
    out, err := strconv.Atoi(in)
    if err != nil {
        fmt.Println(err)
    }
    return out
}
