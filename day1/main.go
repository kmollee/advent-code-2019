package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const inputPath = "input.txt"

// calculateFuel: find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}

func main() {
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int
	var mass int
	for scanner.Scan() {
		mass, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += calculateFuel(mass)
	}

	fmt.Println(total)

}
