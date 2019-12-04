package main

import "testing"

func BenchmarkReadfile0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read0(inputPath)

	}
}

func BenchmarkReadfile1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read1(inputPath)

	}
}

func BenchmarkReadfile2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read2(inputPath)

	}
}
func BenchmarkReadfile3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read3(inputPath)

	}
}
