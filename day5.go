package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	opcode := readCSV()

	newOpcode, output := calcOutput(opcode, 1)
	
	fmt.Println(newOpcode)
	fmt.Println(output)
}

func calcOutput(origOpcode []int64, input int64) ([]int64, []int64){
	opcode := make([]int64, len(origOpcode))
	copy(opcode, origOpcode)
	var outputs []int64
	i := 0
	for i < len(opcode) - 3 {
		fmt.Println(i)
		oc := opcode[i]
		fmt.Println(oc)
		var param1 int64
		var param2 int64
		if oc < 4 {
			param1 = opcode[opcode[i + 1]]
			param2 = opcode[opcode[i + 2]]
		} else if oc == 4 {
			param1 = opcode[opcode[i + 1]]
		} else {
			firstMode := (oc / 100) % 10
			secondMode := (oc / 1000) % 10
			oc = oc % 10
			if firstMode == 0 {
				param1 = opcode[opcode[i + 1]]
			} else {
				param1 = opcode[i + 1]
			}

			if secondMode == 0 {
				param2 = opcode[opcode[i + 2]]
			} else {
				param2 = opcode[i + 2]
			}
		}
		
		if (oc == 99) {
			i = len(opcode)
		} else if (oc == 1) {
			opcode[opcode[i+3]] = param1 + param2
			i += 4
		} else if (oc == 2) {
			opcode[opcode[i+3]] = param1 * param2
			i += 4
		} else if (oc == 3) {
			opcode[opcode[i+1]] = input
			i += 2
		} else if (oc == 4) {
			outputs = append(outputs, param1)
			i += 2
		}
	}

	return opcode, outputs
}

func readCSV() []int64 {
	//read csv of opcode
	file, err := os.Open("./inputs/day5.csv")
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

	return opcode
}