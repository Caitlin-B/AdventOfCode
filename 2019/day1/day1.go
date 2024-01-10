package main

import (
    "encoding/csv"
    "fmt"
    "go.uber.org/zap"
    "math"
    "os"
    "strconv"
)

func main() {
    log := zap.NewExample()

    //read csv of module mass
    file, err := os.Open("./2019/day1/day1.csv")
    if err != nil {
        log.Error("opening csv", zap.Error(err))
        return
    }
    r := csv.NewReader(file)
    modules, err := r.ReadAll()
    if err != nil {
        log.Error("parsing csv", zap.Error(err))
        return
    }

    //calculate fuel requirements

    for _, mass := range modules {
        if s, err := strconv.ParseFloat(mass[0], 32); err == nil {
            calculateFuel(s)
        }
    }

    fmt.Println(sum)
}

var sum float64 = 0

func calculateFuel(mass float64) float64 {
    fuelNeeded := math.Floor(mass/3) - 2

    if fuelNeeded >= 0 {
        sum += fuelNeeded
        calculateFuel(fuelNeeded)
    }

    return sum
}
