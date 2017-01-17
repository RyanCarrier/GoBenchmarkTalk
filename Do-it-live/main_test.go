package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func benchmark(b *testing.B, which int, inputSize int) {
	filename := "../CreateCases/test" + strconv.Itoa(inputSize) + ".in"
	f, _ := os.Open(filename)
	defer f.Close()
	fmt.Fscan(f, &tests)
	input, listInput := get(f)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch which {
		case 0:
			solveA(input, listInput)
		case 1:
			solve2A(input, listInput)
		case 2:
			solveB(input, listInput)
		case 3:
			solveC(input, listInput)
		case 4:
			solve2C(input, listInput)
		}
	}
}

func BenchmarkA10(b *testing.B)     { benchmark(b, 0, 10) }
func BenchmarkA100(b *testing.B)    { benchmark(b, 0, 100) }
func BenchmarkA1000(b *testing.B)   { benchmark(b, 0, 1000) }
func BenchmarkA10000(b *testing.B)  { benchmark(b, 0, 10000) }
func BenchmarkA100000(b *testing.B) { benchmark(b, 0, 100000) }

// func BenchmarkA1000000(b *testing.B)  { benchmark(b, 0, 1000000) }
// func BenchmarkA10000000(b *testing.B) { benchmark(b, 0, 10000000) }

func BenchmarkA210(b *testing.B)     { benchmark(b, 1, 10) }
func BenchmarkA2100(b *testing.B)    { benchmark(b, 1, 100) }
func BenchmarkA21000(b *testing.B)   { benchmark(b, 1, 1000) }
func BenchmarkA210000(b *testing.B)  { benchmark(b, 1, 10000) }
func BenchmarkA2100000(b *testing.B) { benchmark(b, 01, 100000) }

// func BenchmarkA21000000(b *testing.B)  { benchmark(b, 01, 1000000) }
// func BenchmarkA210000000(b *testing.B) { benchmark(b, 01, 10000000) }

func BenchmarkB10(b *testing.B)       { benchmark(b, 02, 10) }
func BenchmarkB100(b *testing.B)      { benchmark(b, 02, 100) }
func BenchmarkB1000(b *testing.B)     { benchmark(b, 02, 1000) }
func BenchmarkB10000(b *testing.B)    { benchmark(b, 02, 10000) }
func BenchmarkB100000(b *testing.B)   { benchmark(b, 02, 100000) }
func BenchmarkB1000000(b *testing.B)  { benchmark(b, 02, 1000000) }
func BenchmarkB10000000(b *testing.B) { benchmark(b, 02, 10000000) }

func BenchmarkC10(b *testing.B)       { benchmark(b, 03, 10) }
func BenchmarkC100(b *testing.B)      { benchmark(b, 03, 100) }
func BenchmarkC1000(b *testing.B)     { benchmark(b, 03, 1000) }
func BenchmarkC10000(b *testing.B)    { benchmark(b, 03, 10000) }
func BenchmarkC100000(b *testing.B)   { benchmark(b, 03, 100000) }
func BenchmarkC1000000(b *testing.B)  { benchmark(b, 03, 1000000) }
func BenchmarkC10000000(b *testing.B) { benchmark(b, 03, 10000000) }

func BenchmarkC210(b *testing.B)       { benchmark(b, 04, 10) }
func BenchmarkC2100(b *testing.B)      { benchmark(b, 04, 100) }
func BenchmarkC21000(b *testing.B)     { benchmark(b, 04, 1000) }
func BenchmarkC210000(b *testing.B)    { benchmark(b, 04, 10000) }
func BenchmarkC2100000(b *testing.B)   { benchmark(b, 04, 100000) }
func BenchmarkC21000000(b *testing.B)  { benchmark(b, 04, 1000000) }
func BenchmarkC210000000(b *testing.B) { benchmark(b, 04, 10000000) }
