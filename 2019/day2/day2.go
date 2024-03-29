package main

import (
    "encoding/csv"
    "fmt"
    "go.uber.org/zap"
    "os"
    "strconv"
)

func main() {
    log := zap.NewExample()

    //read csv of opcode
    file, err := os.Open("./2019/day2/day2.csv")
    if err != nil {
        log.Error("opening csv", zap.Error(err))
        return
    }
    r := csv.NewReader(file)
    opcodearr, err := r.ReadAll()
    opcodestr := opcodearr[0]
    if err != nil {
        log.Error("opening csv", zap.Error(err))
        return
    }

    //opcode str to float64

    opcode := make([]int64, len(opcodestr))
    for i, v := range opcodestr {
        opcode[i], err = strconv.ParseInt(v, 10, 32)
    }
    //try all combinations of verb and noun until one produces required output

    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            testOpcode := make([]int64, len(opcode))
            copy(testOpcode, opcode)
            testOpcode[1] = int64(i)
            testOpcode[2] = int64(j)
            output := calcOutputday2(testOpcode)
            if output[0] == 19690720 {
                fmt.Printf("Noun: %v\nVerb: %v\n", i, j)
            }
        }
    }
}

func calcOutputday2(origOpcode []int64) []int64 {
    opcode := make([]int64, len(origOpcode))
    copy(opcode, origOpcode)
    for i := 0; i < len(opcode)-3; i += 4 {
        param1 := opcode[i+1]
        param2 := opcode[i+2]
        param3 := opcode[i+3]

        if opcode[i] == 99 {
            i = len(opcode)
        } else if opcode[i] == 1 {
            opcode[param3] = opcode[param1] + opcode[param2]
        } else if opcode[i] == 2 {
            opcode[param3] = opcode[param1] * opcode[param2]
        }
    }

    return opcode
}
