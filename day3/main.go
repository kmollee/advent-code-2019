package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

const inputPath = "input.txt"

var directions = map[byte]Point{
	'U': {0, 1},
	'D': {0, -1},
	'L': {-1, 0},
	'R': {1, 0},
}

func main() {
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(findClosestIntersction(string(b)))
	fmt.Println(findClosestIntersctionOnLine(string(b)))
}

type Point struct {
	x, y int
}

type Line []Point

func (l *Line) addPoint(p Point) {
	*l = append(*l, p)
}

func (l *Line) calculateDistance(point Point) (distance int) {
	for i, p := range *l {
		if p == point {
			return i + 1
		}
	}
	return math.MaxInt32
}

func (l *Line) intersect(another Line) []Point {
	seem := make(map[Point]struct{})
	var insertectPoints []Point
	for _, point := range *l {
		seem[point] = struct{}{}
	}

	for _, point := range another {
		if _, exist := seem[point]; exist {
			insertectPoints = append(insertectPoints, point)
		}
	}
	return insertectPoints
}

func route(instructions []string) Line {
	var direction byte
	var distance int

	current := Point{0, 0}
	var wire Line
	for _, ins := range instructions {
		_, err := fmt.Sscanf(ins, "%c%d", &direction, &distance)
		check(err)

		for i := 0; i < distance; i++ {
			current.x += directions[direction].x
			current.y += directions[direction].y
			wire.addPoint(current)
		}
	}
	return wire
}

func calculateManhattanDistance(c Point) int {
	return abs(c.x) + abs(c.y)
}

func findClosestIntersction(wireStr string) int {
	wires := strings.Split(wireStr, "\n")
	if len(wires) != 2 {
		panic("should only have two lines")
	}

	lines := make([]Line, len(wires))
	for i, wire := range wires {
		instructions := strings.Split(wire, ",")
		lines[i] = route(instructions)
	}

	intersectionPoints := lines[0].intersect(lines[1])
	if len(intersectionPoints) == 0 {
		panic("no intersect")
	}

	minManhattanDistance := calculateManhattanDistance(intersectionPoints[0])

	for _, point := range intersectionPoints[1:] {
		minManhattanDistance = min(minManhattanDistance, calculateManhattanDistance(point))
	}

	return minManhattanDistance
}

func findClosestIntersctionOnLine(wireStr string) int {
	wires := strings.Split(wireStr, "\n")
	if len(wires) != 2 {
		panic("should only have two lines")
	}

	lines := make([]Line, len(wires))
	for i, wire := range wires {
		instructions := strings.Split(wire, ",")
		lines[i] = route(instructions)
	}

	intersectionPoints := lines[0].intersect(lines[1])
	if len(intersectionPoints) == 0 {
		panic("no intersect")
	}

	var stepCounts []int
	// now go through each intersection and find the one with the least number of steps
	for _, point := range intersectionPoints {
		stepsWire1 := lines[0].calculateDistance(point)
		stepsWire2 := lines[1].calculateDistance(point)
		stepCounts = append(stepCounts, stepsWire1+stepsWire2)
	}

	sort.Ints(stepCounts)
	return stepCounts[0]
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func panicf(fmtStr string, args ...interface{}) {
	panic(fmt.Sprintf(fmtStr, args...))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
