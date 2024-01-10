package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    strOrbits := readCSV()
    count := 0

    orbits := make([][]string, len(strOrbits))

    for i, v := range strOrbits {
        orbits[i] = append(orbits[i], v[0][0:3], v[0][4:])
    }

    for _, v := range orbits {
        object := v[0]
        j := 0
        for j < len(orbits) {
            for i := 0; i < len(orbits); i++ {
                if orbits[i][1] == object {
                    count += 1
                    object = orbits[i][0]
                    if object == "COM" {
                        j = len(orbits)
                    }
                }
            }
            j += 1
        }
    }

    fmt.Println(count + len(orbits))
    //YOU is at i = 1447
    //SAN is at i = 1493
    youOrbits := orbitsToCOM(orbits, "YOU")
    sanOrbits := orbitsToCOM(orbits, "SAN")

    for i, v1 := range youOrbits {
        for j, v2 := range sanOrbits {
            if v1 == v2 {
                fmt.Println(v1, i+j)
            }
        }
    }

}
func orbitsToCOM(orbits [][]string, object string) []string {
    var transfers []string
    j := 0
    for j < len(orbits) {
        for i := 0; i < len(orbits); i++ {
            if orbits[i][1] == object {
                object = orbits[i][0]
                transfers = append(transfers, object)
                if object == "COM" {
                    j = len(orbits)
                }
            }
        }
        j += 1
    }
    return transfers
}

func readCSV() [][]string {
    //read csv of opcode
    file, err := os.Open("./2019/day6/day6.csv")
    if err != nil {
        fmt.Println("Unable to open csv")
    }
    r := csv.NewReader(file)
    orbits, err := r.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse csv")
    }

    return orbits
}
