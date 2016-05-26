package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var testinput = "test.in"
var biginput = "bigtest.in"
var solver = 2

func benchmark(b *testing.B, which int, n int) {
	filename := "../CreateCases/test" + strconv.Itoa(n) + ".in"
	r, _ := os.Open(filename)
	var tests int
	fmt.Fscan(r, &tests)
	b.ResetTimer()
	for i := 0; i < tests; i++ {
		switch which {
		case 0:
			solveA(get(r))
		case 1:
			solveB(get(r))
		case 2:
			solveC(get(r))
		}
	}
}

func BenchmarkSolve10(b *testing.B) {
	benchmark(b, solver, 10)
}

func BenchmarkSolve100(b *testing.B) {
	benchmark(b, solver, 100)
}
func BenchmarkSolve1000(b *testing.B) {
	benchmark(b, solver, 1000)
}
func BenchmarkSolve10000(b *testing.B) {
	benchmark(b, solver, 10000)
}
func BenchmarkSolve100000(b *testing.B) {
	benchmark(b, solver, 100000)
}
func BenchmarkSolve1000000(b *testing.B) {
	benchmark(b, solver, 1000000)
}
func BenchmarkSolve10000000(b *testing.B) {
	benchmark(b, solver, 10000000)
}
