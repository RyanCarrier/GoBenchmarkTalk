package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

var testinput = "test.in"
var biginput = "bigtest.in"
var part1disable = true

func benchmark(b *testing.B, which int, n int) {

	filename := "../CreateCases/test" + strconv.Itoa(n) + ".in"
	r, _ := os.Open(filename)
	var tests int
	fmt.Fscan(r, &tests)
	inputs, inputs2 := get(r)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		switch which {
		case 0:
			solveA(inputs, inputs2)
		case 1:
			solveB(inputs, inputs2)
		case 2:
			solveC(inputs, inputs2)
		case 3:
			solveA2(inputs, inputs2)
		}

	}
}

func benchmarkGet(b *testing.B, n int) {
	filename := "../CreateCases/test" + strconv.Itoa(n) + ".in"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r, _ := os.Open(filename)
		var tests int
		fmt.Fscan(r, &tests)
		for i := 0; i < tests; i++ {
			get(r)
		}
	}
}

func BenchmarkSolveA10(b *testing.B)       { benchmark(b, 0, 10) }
func BenchmarkSolveA100(b *testing.B)      { benchmark(b, 0, 100) }
func BenchmarkSolveA1000(b *testing.B)     { benchmark(b, 0, 1000) }
func BenchmarkSolveA10000(b *testing.B)    { benchmark(b, 0, 10000) }
func BenchmarkSolveA100000(b *testing.B)   { benchmark(b, 0, 100000) }
func BenchmarkSolveA1000000(b *testing.B)  { benchmark(b, 0, 1000000) }
func BenchmarkSolveA10000000(b *testing.B) { benchmark(b, 0, 10000000) }

func BenchmarkSolve2A10(b *testing.B)       { benchmark(b, 3, 10) }
func BenchmarkSolve2A100(b *testing.B)      { benchmark(b, 3, 100) }
func BenchmarkSolve2A1000(b *testing.B)     { benchmark(b, 3, 1000) }
func BenchmarkSolve2A10000(b *testing.B)    { benchmark(b, 3, 10000) }
func BenchmarkSolve2A100000(b *testing.B)   { benchmark(b, 3, 100000) }
func BenchmarkSolve2A1000000(b *testing.B)  { benchmark(b, 3, 1000000) }
func BenchmarkSolve2A10000000(b *testing.B) { benchmark(b, 3, 10000000) }

func BenchmarkSolveB10(b *testing.B)       { benchmark(b, 1, 10) }
func BenchmarkSolveB100(b *testing.B)      { benchmark(b, 1, 100) }
func BenchmarkSolveB1000(b *testing.B)     { benchmark(b, 1, 1000) }
func BenchmarkSolveB10000(b *testing.B)    { benchmark(b, 1, 10000) }
func BenchmarkSolveB100000(b *testing.B)   { benchmark(b, 1, 100000) }
func BenchmarkSolveB1000000(b *testing.B)  { benchmark(b, 1, 1000000) }
func BenchmarkSolveB10000000(b *testing.B) { benchmark(b, 1, 10000000) }

func BenchmarkSolveC10(b *testing.B)       { benchmark(b, 2, 10) }
func BenchmarkSolveC100(b *testing.B)      { benchmark(b, 2, 100) }
func BenchmarkSolveC1000(b *testing.B)     { benchmark(b, 2, 1000) }
func BenchmarkSolveC10000(b *testing.B)    { benchmark(b, 2, 10000) }
func BenchmarkSolveC100000(b *testing.B)   { benchmark(b, 2, 100000) }
func BenchmarkSolveC1000000(b *testing.B)  { benchmark(b, 2, 1000000) }
func BenchmarkSolveC10000000(b *testing.B) { benchmark(b, 2, 10000000) }

/*
func BenchmarkGet10(b *testing.B)         { benchmarkGet(b, 10) }
func BenchmarkGet100(b *testing.B)        { benchmarkGet(b, 100) }
func BenchmarkGet1000(b *testing.B)       { benchmarkGet(b, 1000) }
func BenchmarkGet10000(b *testing.B)      { benchmarkGet(b, 10000) }
func BenchmarkGet100000(b *testing.B)     { benchmarkGet(b, 100000) }
func BenchmarkGet1000000(b *testing.B)    { benchmarkGet(b, 1000000) }
*/

func test(t *testing.T, which int, n int) {
	filename := "../CreateCases/test" + strconv.Itoa(n) + ".in"
	filenameout := "../CreateCases/test" + strconv.Itoa(n) + ".out"
	r, _ := os.Open(filename)
	var tests int
	fmt.Fscan(r, &tests)
	inputs, inputs2 := get(r)
	body, _ := ioutil.ReadFile(filenameout)
	parts := strings.Fields(string(body))
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	var solA, solB int
	switch which {
	case 0:
		solA, solB = solveA(inputs, inputs2)
	case 1:
		solA, solB = solveB(inputs, inputs2)
	case 2:
		solA, solB = solveC(inputs, inputs2)
	case 3:
		solA, solB = solveA2(inputs, inputs2)
	}
	switch {
	case solA != a && solB != b:
		if solA != b && solB != a {
			t.Error("\nA GOT:", solA, "\nA WANT:", a, "\nB GOT:", solB, "\nB WANT:", b)
		}
	case solA != a:
		t.Error("\nA GOT:", solA, "\nA WANT:", a)
	case solB != b:
		t.Error("\nB GOT:", solB, "\nB WANT:", b)
	}
}

func TestSolveA10(t *testing.T)     { test(t, 0, 10) }
func TestSolveA100(t *testing.T)    { test(t, 0, 100) }
func TestSolveA1000(t *testing.T)   { test(t, 0, 1000) }
func TestSolveA10000(t *testing.T)  { test(t, 0, 10000) }
func TestSolveA100000(t *testing.T) { test(t, 0, 100000) }

func TestSolve2A10(t *testing.T)     { test(t, 3, 10) }
func TestSolve2A100(t *testing.T)    { test(t, 3, 100) }
func TestSolve2A1000(t *testing.T)   { test(t, 3, 1000) }
func TestSolve2A10000(t *testing.T)  { test(t, 3, 10000) }
func TestSolve2A100000(t *testing.T) { test(t, 3, 100000) }

func TestSolveB10(t *testing.T)     { test(t, 1, 10) }
func TestSolveB100(t *testing.T)    { test(t, 1, 100) }
func TestSolveB1000(t *testing.T)   { test(t, 1, 1000) }
func TestSolveB10000(t *testing.T)  { test(t, 1, 10000) }
func TestSolveB100000(t *testing.T) { test(t, 1, 100000) }

func TestSolveC10(t *testing.T)     { test(t, 2, 10) }
func TestSolveC100(t *testing.T)    { test(t, 2, 100) }
func TestSolveC1000(t *testing.T)   { test(t, 2, 1000) }
func TestSolveC10000(t *testing.T)  { test(t, 2, 10000) }
func TestSolveC100000(t *testing.T) { test(t, 2, 100000) }
