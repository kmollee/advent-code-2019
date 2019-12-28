package main

import (
	"fmt"
	"log"

	"strings"

	"github.com/kmollee/advent-code-2019/pkg/aocutils"
)

const inputPath = "input.txt"

type parameterMode int
type opcode int

const (
	position  parameterMode = 0
	immediate parameterMode = 1
)
const (
	add    opcode = 1
	mul    opcode = 2
	input  opcode = 3
	output opcode = 4
	stop   opcode = 99

	/* part2 extend opcode */
	jumpIfTrue  opcode = 5
	jumpIfFalse opcode = 6
	lessThan    opcode = 7
	equals      opcode = 8
)

func main() {
	fmt.Println("========part 1===========")
	part1()

	fmt.Println("========part 2===========")
	part2()
}

func part1() {
	inputs := loadInputs(inputPath)
	userInput := 1
	instructionIndex := 0
loop:
	for {
		op, paramsMode := parseInstruction(inputs[instructionIndex])
		switch op {
		case add:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			writeValue(inputs, instructionIndex+3, v1+v2)
			instructionIndex += 4
		case mul:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			writeValue(inputs, instructionIndex+3, v1*v2)
			instructionIndex += 4

		case input:
			writeValue(inputs, instructionIndex+1, userInput)
			instructionIndex += 2

		case output:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			fmt.Println(v1)
			instructionIndex += 2
		case stop:
			break loop

		default:
			log.Fatalf("Unknow opcode:%v", op)
		}
	}
}

func part2() {
	inputs := loadInputs(inputPath)
	userInput := 5
	instructionIndex := 0
loop:
	for {
		op, paramsMode := parseInstruction(inputs[instructionIndex])
		// log.Println(inputs[instructionIndex], op, paramsMode)
		switch op {
		case add:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			writeValue(inputs, instructionIndex+3, v1+v2)
			instructionIndex += 4
		case mul:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			writeValue(inputs, instructionIndex+3, v1*v2)
			instructionIndex += 4

		case input:
			writeValue(inputs, instructionIndex+1, userInput)
			instructionIndex += 2
		case output:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			fmt.Println(v1)
			instructionIndex += 2

		case stop:
			break loop

		case jumpIfTrue:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			if v1 != 0 {
				v2 := readValue(inputs, instructionIndex, 2, paramsMode)
				instructionIndex = v2
			} else {
				instructionIndex += 3
			}
		case jumpIfFalse:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			if v1 == 0 {
				v2 := readValue(inputs, instructionIndex, 2, paramsMode)
				instructionIndex = v2
			} else {
				instructionIndex += 3
			}
		case lessThan:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			if v1 < v2 {
				writeValue(inputs, instructionIndex+3, 1)
			} else {
				writeValue(inputs, instructionIndex+3, 0)
			}
			instructionIndex += 4
		case equals:
			v1 := readValue(inputs, instructionIndex, 1, paramsMode)
			v2 := readValue(inputs, instructionIndex, 2, paramsMode)
			if v1 == v2 {
				writeValue(inputs, instructionIndex+3, 1)
			} else {
				writeValue(inputs, instructionIndex+3, 0)
			}
			instructionIndex += 4
		default:
			log.Fatalf("Unknown opcode: %d from instruction at index: %d\n", op, instructionIndex)
		}
	}
}

func loadInputs(path string) []int {
	b := aocutils.LoadFile(path)
	data := strings.Split(string(b), ",")

	var inputs []int
	for _, valueStr := range data {
		val := aocutils.Atoi(valueStr)
		inputs = append(inputs, val)
	}
	return inputs
}

/*
ABCDE
 1002

DE - two-digit opcode,      02 == opcode 2
 C - mode of 1st parameter,  0 == position mode
 B - mode of 2nd parameter,  1 == immediate mode
 A - mode of 3rd parameter,  0 == position mode,
                                  omitted due to being a leading zero
*/
func parseInstruction(instruction int) (opcode, []parameterMode) {
	var paramsMode []parameterMode

	opcode := opcode(instruction % 100)
	instruction /= 100

	for instruction > 0 {
		paramMode := parameterMode(instruction % 10)
		paramsMode = append(paramsMode, paramMode)

		instruction /= 10
	}

	return opcode, paramsMode
}

func readValue(input []int, instructionIndex, parameterIndex int, parametersMode []parameterMode) int {

	paramMode := position
	if len(parametersMode) > parameterIndex-1 {
		paramMode = parametersMode[parameterIndex-1]
	}

	var output int

	switch paramMode {
	case position:
		readIndex := input[instructionIndex+parameterIndex]
		output = input[readIndex]
	case immediate:
		output = input[instructionIndex+parameterIndex]
	default:
		log.Fatalf("Unknown parameter mode: %d from instruction at index: %d\n", paramMode, instructionIndex)
	}

	return output
}

func writeValue(input []int, index, value int) {
	readIndex := input[index]
	input[readIndex] = value
}
