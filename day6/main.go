package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kmollee/advent-code-2019/pkg/aocutils"
)

func main() {
	fmt.Println("========part 1===========")
	part1()

	fmt.Println("========part 2===========")
	part2()
}

func part1() {
	b := aocutils.LoadFile("input.txt")
	buf := bytes.NewBuffer(b)
	scanner := bufio.NewScanner(buf)

	orbitmap := newOrbitMap()
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		center, satellite := s[0], s[1]
		orbitmap.add(center, satellite)
	}

	total := 0
	for _, orbit := range orbitmap {
		total += orbit.checksum()
	}
	log.Println(total)
}

func part2() {
	b := aocutils.LoadFile("input.txt")
	buf := bytes.NewBuffer(b)
	scanner := bufio.NewScanner(buf)

	orbitmap := newOrbitMap()
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ")")
		center, satellite := s[0], s[1]
		orbitmap.add(center, satellite)

	}

	you := orbitmap.createOrGetOrbit("YOU")
	san := orbitmap.createOrGetOrbit("SAN")
	orbits := commonOrbits(you.getAllParents(), san.getAllParents())
	if len(orbits) == 0 {
		log.Fatal("no common Orbit")
	}

	lastCommonOrbit := orbits[0]

	// is parent orbital transfers, not include self
	// need minus 1
	yd, err := you.getDistance(lastCommonOrbit)
	aocutils.Check(err)
	yd = yd - 1

	sd, err := san.getDistance(lastCommonOrbit)
	aocutils.Check(err)
	sd = sd - 1

	log.Println(yd + sd)

}

type orbit struct {
	name string
	prev *orbit
	next []*orbit
}

func newOrbit(name string) *orbit {
	return &orbit{name: name}
}

func (o *orbit) checksum() int {
	count := 0
	for current := o; current.prev != nil; current = current.prev {
		count++
	}
	return count
}

func (o *orbit) getAllParents() []*orbit {

	var orbits []*orbit

	for current := o; current != nil; current = current.prev {
		orbits = append(orbits, current)
	}
	return orbits
}

func (o *orbit) add(other *orbit) {
	other.prev = o
	o.next = append(o.next, other)
}

func (o *orbit) getDistance(parent *orbit) (int, error) {
	count := 0
	for current := o; current != nil; current = current.prev {
		if current == parent {
			return count, nil
		}
		count++
	}
	return -1, fmt.Errorf("could not found %v", parent)
}

func commonOrbits(a, b []*orbit) []*orbit {
	var orbits []*orbit
	for _, ao := range a {
		for _, bo := range b {
			if ao == bo {
				orbits = append(orbits, bo)
			}
		}
	}
	return orbits

}

type OrbitMap map[string]*orbit

func (o OrbitMap) add(center, satellite string) {
	c := o.createOrGetOrbit(center)
	s := o.createOrGetOrbit(satellite)
	c.add(s)
}

func (o OrbitMap) getOrbit(name string) *orbit {
	if orbit, exist := o[name]; exist {
		return orbit
	}
	return nil
}

func (o OrbitMap) createOrGetOrbit(name string) *orbit {
	if orbit, exist := o[name]; exist {
		return orbit
	}
	o[name] = newOrbit(name)
	return o[name]
}

func newOrbitMap() OrbitMap {
	return OrbitMap(make(map[string]*orbit))
}
