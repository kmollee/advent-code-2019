package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const inputPath = "input.txt"

func main() {
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(findClosestIntersction(string(b)))
	fmt.Println(findClosestIntersctionOnLine(string(b)))

}

type coordinate struct {
	x, y int
}

type line []coordinate

func (l *line) calculateDistance(point coordinate) int {
	for i, p := range *l {
		if p == point {
			return i + 1
		}
	}
	return math.MaxInt32
}

func (c *coordinate) StepX(distance int) line {
	if distance == 0 {
		return nil
	}

	var history line
	if distance < 0 {
		for i := 0; i > distance; i-- {
			c.x--
			history = append(history, coordinate{c.x, c.y})
		}
	} else {
		for i := 0; i < distance; i++ {
			c.x++
			history = append(history, coordinate{c.x, c.y})
		}
	}
	return history
}

func (c *coordinate) StepY(distance int) line {
	if distance == 0 {
		return nil
	}

	var history line
	if distance < 0 {
		for i := 0; i > distance; i-- {
			c.y--
			history = append(history, coordinate{c.x, c.y})
		}
	} else {
		for i := 0; i < distance; i++ {
			c.y++
			history = append(history, coordinate{c.x, c.y})
		}
	}
	return history
}

func route(instructions []string) line {
	var direction byte
	var distance int

	current := coordinate{0, 0}
	var history line
	for _, ins := range instructions {
		_, err := fmt.Sscanf(ins, "%c%d", &direction, &distance)
		if err != nil {
			panic(err)
		}

		switch direction {
		case 'L':
			history = append(history, current.StepX(-distance)...)
		case 'R':
			history = append(history, current.StepX(distance)...)
		case 'U':
			history = append(history, current.StepY(distance)...)
		case 'D':
			history = append(history, current.StepY(-distance)...)
		default:
			panic(fmt.Sprintf("unexpect instruction:%s", ins))
		}
	}
	return history
}

func calculateManhattanDistance(c coordinate) int {
	return abs(c.x) + abs(c.y)
}

func findIntersection(a, b line) []coordinate {
	seem := make(map[coordinate]struct{})
	var insertectPoints []coordinate
	for _, point := range a {
		seem[point] = struct{}{}
	}

	for _, point := range b {
		if _, exist := seem[point]; exist {
			insertectPoints = append(insertectPoints, point)
		}
	}
	return insertectPoints
}

func findClosestIntersction(wireStr string) int {
	wires := strings.Split(wireStr, "\n")
	// should only have two lines
	if len(wires) != 2 {
		return 0
	}

	lines := make([]line, len(wires))
	for i, wire := range wires {
		instructions := strings.Split(wire, ",")
		lines[i] = route(instructions)
	}

	intersectionPoints := findIntersection(lines[0], lines[1])

	// no interseciton point
	if len(intersectionPoints) == 0 {
		return 0
	}

	minManhattanDistance := calculateManhattanDistance(intersectionPoints[0])

	for _, point := range intersectionPoints[1:] {
		minManhattanDistance = min(minManhattanDistance, calculateManhattanDistance(point))

	}

	return minManhattanDistance
}

func findClosestIntersctionOnLine(wireStr string) int {
	wires := strings.Split(wireStr, "\n")
	// should only have two lines
	if len(wires) != 2 {
		return 0
	}

	lines := make([]line, len(wires))
	for i, wire := range wires {
		instructions := strings.Split(wire, ",")
		lines[i] = route(instructions)
	}

	intersectionPoints := findIntersection(lines[0], lines[1])

	// no interseciton point
	if len(intersectionPoints) == 0 {
		return 0
	}

	minLineDistance := lines[0].calculateDistance(intersectionPoints[0]) + lines[1].calculateDistance(intersectionPoints[0])

	for _, point := range intersectionPoints[1:] {

		minLineDistance = min(minLineDistance, lines[0].calculateDistance(point)+lines[1].calculateDistance(point))
	}

	return minLineDistance
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
