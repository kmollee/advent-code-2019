package aocutils

import (
	"bytes"
	"io/ioutil"
)

func LoadFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	b = bytes.TrimSuffix(b, []byte("\n"))
	Check(err)
	return b
}
