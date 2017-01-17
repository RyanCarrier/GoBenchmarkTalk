package main

import (
	"runtime"
	t "testing"
)

func bench(b *t.B, x int, sqrtCheck bool) {
	for n := 0; n <= b.N; n++ {
		test(x, sqrtCheck)
	}
}

func sbench(b *t.B, x, routines int, quitEnabled bool) {
	if x%routines != 0 {
		b.Errorf("Can't split %d into %d go routines", x, routines)
	}
	for n := 0; n <= b.N; n++ {
		a, _, _ := superTest(x, routines, quitEnabled)
		if a == 0 {
			b.Error("Either goal is unatainable (most likely)")
		}
	}
}

/*
//With sqrtchecking early
func BenchmarkTestSqrt10(b *t.B)     { bench(b, 12, true) }
func BenchmarkTestSqrt100(b *t.B)    { bench(b, 108, true) }
func BenchmarkTestSqrt1000(b *t.B)   { bench(b, 1000, true) }
func BenchmarkTestSqrt10000(b *t.B)  { bench(b, 10000, true) }
func BenchmarkTestSqrt100000(b *t.B) { bench(b, 100000, true) }

//With normal, sqrt checking during solution found
func BenchmarkTestPlain10(b *t.B)     { bench(b, 12, false) }
func BenchmarkTestPlain100(b *t.B)    { bench(b, 108, false) }
func BenchmarkTestPlain1000(b *t.B)   { bench(b, 1000, false) }
func BenchmarkTestPlain10000(b *t.B) { bench(b, 10000, false) }*/
//func BenchmarkTestPlain100000(b *t.B) { bench(b, 100000, false) }

//With multiple gorountines, which all quit when any routine finds a solution
//func BenchmarkSuperTest10000Q4(b *t.B)    { sbench(b, 100000, 4, true) }
//func BenchmarkSuperTest10000Q16(b *t.B)   { sbench(b, 100000, 16, true) }
func BenchmarkSuperTest10000Q_CPU0_5(b *t.B) { sbench(b, 100000, runtime.NumCPU()/2, true) }
func BenchmarkSuperTest10000Q_CPU1(b *t.B)   { sbench(b, 100000, runtime.NumCPU()*1, true) }
func BenchmarkSuperTest10000Q_CPU2(b *t.B)   { sbench(b, 100000, runtime.NumCPU()*2, true) }
func BenchmarkSuperTest10000Q_CPU4(b *t.B)   { sbench(b, 100000, runtime.NumCPU()*4, true) }
func BenchmarkSuperTest10000Q_CPU8(b *t.B)   { sbench(b, 100000, runtime.NumCPU()*8, true) }
func BenchmarkSuperTest10000Q_CPU16(b *t.B)  { sbench(b, 100000, runtime.NumCPU()*16, true) }

/*
//
func BenchmarkSuperTest10000NQ4(b *t.B)    { sbench(b, 100000, 4, false) }
func BenchmarkSuperTest10000NQ16(b *t.B)   { sbench(b, 100000, 16, false) }
func BenchmarkSuperTest10000NQ_CPU(b *t.B) { sbench(b, 100000, runtime.NumCPU(), false) }
*/
