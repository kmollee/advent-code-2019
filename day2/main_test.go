package main

import (
	"reflect"
	"testing"
)

/*
   For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
   For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
   For a mass of 1969, the fuel required is 654.
   For a mass of 100756, the fuel required is 33583.
*/
func TestIntcode(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		expect []int
	}{
		{desc: "add", input: "1,0,0,0,99", expect: []int{2, 0, 0, 0, 99}},
		{desc: "mul", input: "2,3,0,3,99", expect: []int{2, 3, 0, 6, 99}},
		{desc: "mul", input: "2,4,4,5,99,0", expect: []int{2, 4, 4, 5, 99, 9801}},
		{desc: "add and mul", input: "1,1,1,4,99,5,6,0,99", expect: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Intcode(tC.input)
			if !reflect.DeepEqual(got, tC.expect) {
				t.Errorf("got '%v' expect '%v'", got, tC.expect)
			}
		})
	}
}
