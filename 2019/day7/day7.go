package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func main() {
    opcode := readCSV()

    //first input is phase setting, second is prev output
    fmt.Println(amplifiers(opcode, 9, 8, 7, 6, 5))

    //phaseSequences:= permutations([]int{0, 1, 2, 3, 4})
    //var maxThruster int64 = 0

    // for _, phases := range phaseSequences {
    // 	thruster := amplifiers(opcode, int64(phases[0]), int64(phases[1]), int64(phases[2]), int64(phases[3]), int64(phases[4]))
    // 	if  thruster > maxThruster {
    // 		maxThruster = thruster
    // 	}
    // }

    //fmt.Println(maxThruster)

}

func amplifiers(opcode []int64, amp1, amp2, amp3, amp4, amp5 int64) int64 {
    _, output1 := calcOutput(opcode, amp1, 0)
    _, output2 := calcOutput(opcode, amp2, output1[0])
    _, output3 := calcOutput(opcode, amp3, output2[0])
    _, output4 := calcOutput(opcode, amp4, output3[0])
    _, output5 := calcOutput(opcode, amp5, output4[0])

    return output5[0]
}

func calcOutput(origOpcode []int64, input1 int64, input2 int64) ([]int64, []int64) {
    opcode := make([]int64, len(origOpcode))
    copy(opcode, origOpcode)
    var outputs []int64
    input := 1
    i := 0
    for i < len(opcode)-3 {
        oc := opcode[i]
        var param1 int64
        var param2 int64
        if oc == 4 || oc == 3 {
            param1 = opcode[opcode[i+1]]
        } else if oc < 10 {
            param1 = opcode[opcode[i+1]]
            param2 = opcode[opcode[i+2]]
        } else if oc != 99 {
            firstMode := (oc / 100) % 10
            secondMode := (oc / 1000) % 10
            oc = oc % 10
            if firstMode == 0 {
                param1 = opcode[opcode[i+1]]
            } else {
                param1 = opcode[i+1]
            }

            if secondMode == 0 {
                param2 = opcode[opcode[i+2]]
            } else {
                param2 = opcode[i+2]
            }
        }

        if oc == 99 {
            i = len(opcode)
        } else if oc == 1 {
            opcode[opcode[i+3]] = param1 + param2
            i += 4
        } else if oc == 2 {
            opcode[opcode[i+3]] = param1 * param2
            i += 4
        } else if oc == 3 {
            if input == 1 {
                opcode[opcode[i+1]] = input1
                input = 2
            } else if input == 2 {
                opcode[opcode[i+1]] = input2
            }
            i += 2
        } else if oc == 4 {
            outputs = append(outputs, param1)
            i += 2
        } else if oc == 5 {
            if param1 != 0 {
                i = int(param2)
            } else {
                i += 3
            }
        } else if oc == 6 {
            if param1 == 0 {
                i = int(param2)
            } else {
                i += 3
            }
        } else if oc == 7 {
            if param1 < param2 {
                opcode[opcode[i+3]] = 1
            } else {
                opcode[opcode[i+3]] = 0
            }
            i += 4
        } else if oc == 8 {
            if param1 == param2 {
                opcode[opcode[i+3]] = 1
            } else {
                opcode[opcode[i+3]] = 0
            }
            i += 4
        }
    }

    return opcode, outputs
}

func permutations(arr []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int) {
        if n == 1 {
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(arr, n-1)
                if n%2 == 1 {
                    tmp := arr[i]
                    arr[i] = arr[n-1]
                    arr[n-1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n-1]
                    arr[n-1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

func readCSV() []int64 {
    //read csv of opcode
    file, err := os.Open("./2019/day7/day7.csv")
    if err != nil {
        fmt.Println("Unable to open csv")
    }
    r := csv.NewReader(file)
    opcodearr, err := r.ReadAll()
    opcodestr := opcodearr[0]
    if err != nil {
        fmt.Println("Unable to parse csv")
    }

    //opcode str to float64

    opcode := make([]int64, len(opcodestr))
    for i, v := range opcodestr {
        opcode[i], err = strconv.ParseInt(v, 10, 32)
    }

    return opcode
}
