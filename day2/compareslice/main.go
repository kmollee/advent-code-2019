package main

import (
	"fmt"
	"reflect"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	// array can use the comparison operators == and !=.
	a := [2]int{1, 2}
	b := [2]int{1, 3}
	fmt.Println(a == b) // false

	// General-purpose code for recursive comparison
	var c []int = nil
	var d []int = make([]int, 0)
	fmt.Println(reflect.DeepEqual(c, d)) // false

}
