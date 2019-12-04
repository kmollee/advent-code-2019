package main

import "testing"

/*
   For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
   For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
   For a mass of 1969, the fuel required is 654.
   For a mass of 100756, the fuel required is 33583.
*/
func TestCalculateFuel(t *testing.T) {
	testCases := []struct {
		desc   string
		mass   int
		expect int
	}{
		{desc: "mass of 12", mass: 12, expect: 2},
		{desc: "mass of 14", mass: 14, expect: 2},
		{desc: "mass of 1969", mass: 1969, expect: 654},
		{desc: "mass of 100756", mass: 100756, expect: 33583},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculateFuel(tC.mass)
			if got != tC.expect {
				t.Errorf("got '%v' expect '%v'", got, tC.expect)
			}
		})
	}
}

func TestCalculateFuelWithFiction(t *testing.T) {
	testCases := []struct {
		desc   string
		mass   int
		expect int
	}{
		{desc: "mass of 14", mass: 14, expect: 2},
		{desc: "mass of 1969", mass: 1969, expect: 966},
		{desc: "mass of 100756", mass: 100756, expect: 50346},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculateFuelWithFiciton(tC.mass)
			if got != tC.expect {
				t.Errorf("got '%v' expect '%v'", got, tC.expect)
			}
		})
	}
}
