package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/kmollee/advent-code-2019/pkg/aocutils"
	"github.com/kmollee/advent-code-2019/pkg/aocutils/incode"
)

const inputPath = "input.txt"

func main() {
	fmt.Println("========part 1===========")
	part1()

	fmt.Println("========part 2===========")
	part2()
}

func part1() {
	inputs := loadInputs(inputPath)

	phase := []int{0, 1, 2, 3, 4}

	maxValue := 0
	maxSeq := []int{}

	for _, seq := range permutations(phase) {
		value := amplify(inputs, seq)
		if value > maxValue {
			maxValue = value
			maxSeq = seq
		}
	}
	log.Println(maxSeq, maxValue)
}

func part2() {
	inputs := loadInputs(inputPath)

	phase := []int{5, 6, 7, 8, 9}

	maxValue := 0
	maxSeq := []int{}

	for _, seq := range permutations(phase) {
		value := amplify(inputs, seq)
		if value > maxValue {
			maxValue = value
			maxSeq = seq
		}
	}

	log.Println(maxSeq, maxValue)
}

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
// generate all combination of array
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

func amplify(program, phase []int) int {

	o := []chan int{
		make(chan int, 1),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	var wg sync.WaitGroup

	for i, p := range phase {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			programCpy := make([]int, len(program))
			copy(programCpy, program)
			incode.Incode(programCpy, o[i], o[(i+1)%5])
			wg.Done()
		}(i, &wg)
		o[i] <- p
	}

	o[0] <- 0
	wg.Wait()

	return <-o[0]
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
