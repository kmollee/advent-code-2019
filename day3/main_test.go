package main

import (
	"reflect"
	"testing"
)

func TestRoute(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		expect Line
	}{
		{
			desc:   "right",
			input:  []string{"R1", "R1"},
			expect: Line{{1, 0}, {2, 0}},
		},
		{
			desc:   "right and left",
			input:  []string{"L1", "L1", "R1"},
			expect: Line{{-1, 0}, {-2, 0}, {-1, 0}},
		},
		{
			desc:   "up and right",
			input:  []string{"U1", "R1"},
			expect: Line{{0, 1}, {1, 1}},
		},
		{
			desc:   "up right down",
			input:  []string{"U1", "R1", "D1"},
			expect: Line{{0, 1}, {1, 1}, {1, 0}},
		},
		{
			desc:   "circle",
			input:  []string{"U1", "R1", "D1", "L1"},
			expect: Line{{0, 1}, {1, 1}, {1, 0}, {0, 0}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				got := route(tC.input)
				if !reflect.DeepEqual(got, tC.expect) {
					t.Errorf("got '%v' expect '%v'", got, tC.expect)
				}
			})
		})
	}
}

func TestFindClosestIntersection(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		expect int
	}{
		{
			desc:   "",
			input:  "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
			expect: 159,
		},
		{
			desc:   "",
			input:  "R8,U5,L5,D3\nU7,R6,D4,L4",
			expect: 6,
		},
		{
			desc:   "",
			input:  "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expect: 135,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := findClosestIntersction(tC.input)
			if got != tC.expect {
				t.Errorf("got '%v' expect '%v'", got, tC.expect)
			}
		})
	}
}
