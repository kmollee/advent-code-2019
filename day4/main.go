package main

import (
	"sort"
	"fmt"
	"reflect"
	"strconv"
)

type adjustFunc func(string)bool

func adjust1(num string) bool {
	seem := make(map[byte]struct{})
	for _, n := range []byte(num) {
		if _, exist := seem[n]; exist {
			return true
		}
		seem[n] = struct{}{}
	}
	return false
}


func adjust2(num string) bool {
	seem := make(map[byte]int)
	for _, n := range []byte(num) {
		seem[n]++
	}

	for _, val := range seem {
		if (val == 2){
			return true
		}
	}
	return false
}

func isIncreaseNum(num string) bool {

	n := []byte(num)
	sort.Slice(n, func(i int, j int) bool { return num[i] < num[j] })
	return reflect.DeepEqual([]byte(num), n)
}

func countPuzzleCombination(begin, end int, adjust adjustFunc) int {

	var count int
	for i := begin; i < end; i++ {
		s := strconv.Itoa(i)
		if isIncreaseNum(s) && adjust(s){
			count++
		}
	}
	return count
}


func main() {
	fmt.Println(countPuzzleCombination(246515, 739105, adjust1))
	fmt.Println(countPuzzleCombination(246515, 739105, adjust2))
}
