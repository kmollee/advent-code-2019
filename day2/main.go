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

func convertToInts(s []string) []int {
	n := []int{}
	for _, v := range s {
		m, _ := strconv.Atoi(v)
		n = append(n, m)
	}
	return n
}

func intcode(nums []int, target int) bool {

	i := 0
loop:
	for i < len(nums) {
		switch nums[i] {
		case Add:
			firstIndex := nums[i+1]
			secondIndex := nums[i+2]
			storeIndex := nums[i+3]
			nums[storeIndex] = nums[firstIndex] + nums[secondIndex]
			i += 4
		case Mul:
			firstIndex := nums[i+1]
			secondIndex := nums[i+2]
			storeIndex := nums[i+3]
			nums[storeIndex] = nums[firstIndex] * nums[secondIndex]
			i += 4
		case Halt:
			break loop
		default:
			i++
		}
		if nums[0] == target {
			return true
		}
	}
	return false
}

func part1(list string) int {

	numStrs := strings.Split(list, ",")
	nums := convertToInts(numStrs)

	intcode(nums, -1)

	return nums[0]
}

func part2(list string, target int) (int, int) {

	numStrs := strings.Split(list, ",")
	nums := convertToInts(numStrs)

	for ni := 0; ni < 100; ni++ {
		for vi := 0; vi < 100; vi++ {
			// reset the computer's memory to the values in the program
			solution := make([]int, len(nums))
			copy(solution, nums)
			solution[1] = ni
			solution[2] = vi
			if intcode(solution, target) {
				return ni, vi
			}
		}
	}
	return 0, 0
}

func main() {
	// / Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	// part1
	p1ans := part1(string(b))
	fmt.Println(p1ans)

	// part2
	noun, verb := part2(string(b), 19690720)
	fmt.Println(noun*100 + verb)

}
