package incode

import (
	"log"
)

type parameterMode int
type Opcode int

const (
	position  parameterMode = 0
	immediate parameterMode = 1
)
const (
	Add         Opcode = 1
	Mul         Opcode = 2
	Input       Opcode = 3
	Output      Opcode = 4
	JumpIfTrue  Opcode = 5
	JumpIfFalse Opcode = 6
	LessThan    Opcode = 7
	Equals      Opcode = 8
	Stop        Opcode = 99
)

func Incode(program []int, input <-chan int, output chan<- int) {
	instructionIndex := 0

	for {
		op, paramsMode := parseInstruction(program[instructionIndex])
		switch op {
		case Add:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			v2 := readValue(program, instructionIndex, 2, paramsMode)
			writeValue(program, instructionIndex+3, v1+v2)
			instructionIndex += 4
		case Mul:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			v2 := readValue(program, instructionIndex, 2, paramsMode)
			writeValue(program, instructionIndex+3, v1*v2)
			instructionIndex += 4

		case Input:
			val := <-input
			writeValue(program, instructionIndex+1, val)
			instructionIndex += 2
		case Output:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			output <- v1
			instructionIndex += 2
		case Stop:
			return

		case JumpIfTrue:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			if v1 != 0 {
				v2 := readValue(program, instructionIndex, 2, paramsMode)
				instructionIndex = v2
			} else {
				instructionIndex += 3
			}
		case JumpIfFalse:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			if v1 == 0 {
				v2 := readValue(program, instructionIndex, 2, paramsMode)
				instructionIndex = v2
			} else {
				instructionIndex += 3
			}
		case LessThan:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			v2 := readValue(program, instructionIndex, 2, paramsMode)
			if v1 < v2 {
				writeValue(program, instructionIndex+3, 1)
			} else {
				writeValue(program, instructionIndex+3, 0)
			}
			instructionIndex += 4
		case Equals:
			v1 := readValue(program, instructionIndex, 1, paramsMode)
			v2 := readValue(program, instructionIndex, 2, paramsMode)
			if v1 == v2 {
				writeValue(program, instructionIndex+3, 1)
			} else {
				writeValue(program, instructionIndex+3, 0)
			}
			instructionIndex += 4
		default:
			log.Fatalf("Unknown opcode: %d from instruction at index: %d\n", op, instructionIndex)
		}
	}

	return
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
func parseInstruction(instruction int) (Opcode, []parameterMode) {
	var paramsMode []parameterMode

	opcode := Opcode(instruction % 100)
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
