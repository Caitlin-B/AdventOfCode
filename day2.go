package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//read csv of opcode
	file, err := os.Open("./inputs/day2.csv")
	if err != nil {
		fmt.Println("Unable to open csv") 
	}
	r:= csv.NewReader(file)
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
	fmt.Println(opcode)

	//loop through opcode + calc
	for i := 0; i < len(opcode); i += 4 {
		if (opcode[i] == 99) {
			i = len(opcode)
		} else if (opcode[i] == 1) {
			opcode[opcode[i + 3]] = opcode[opcode[i + 1]] + opcode[opcode[i + 2]]
		} else if (opcode[i] == 2) {
			opcode[opcode[i + 3]] = opcode[opcode[i + 1]] * opcode[opcode[i + 2]]
		}
	}

	fmt.Println(opcode)
}