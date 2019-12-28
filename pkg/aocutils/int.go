package aocutils

import "strconv"

// Atoi gets an int from a string.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}
