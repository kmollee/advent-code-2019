package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	Add  = 1
	Mul  = 2
	Halt = 99
)

const inputPath = "input.txt"

func Intcode(list string) []int {
	// 0, 1, 2
	// other: something went wrong
	numStrs := strings.Split(list, ",")
	nums := make([]int, len(numStrs))
	for i, num := range numStrs {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = n
	}

	i := 0
loop:
	for i < len(nums) {
		switch nums[i] {
		case Add:
			i++
			firstIndex := nums[i]
			i++
			secondIndex := nums[i]
			i++
			storeIndex := nums[i]
			nums[storeIndex] = nums[firstIndex] + nums[secondIndex]
		case Mul:
			i++
			firstIndex := nums[i]
			i++
			secondIndex := nums[i]
			i++
			storeIndex := nums[i]
			nums[storeIndex] = nums[firstIndex] * nums[secondIndex]
		case Halt:
			break loop
		default:
			i++
		}
	}
	return nums
}

func main() {
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}
	fmt.Println(Intcode(string(b)))
}
