package main

import (
	"fmt"
	"math"
	"sync"

	gbd "github.com/RyanCarrier/GoBenchDyn"
)

/*
https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {
	f, ivs := benchWrapper(1000, 1000000, 10, 1, 64, 0, 2, false)
	fmt.Println(gbd.RangeN("SuperBench", f, ivs...))
}

func benchWrapper(xfrom, xto, xmultiple, cpufrom, cputo, cpuinc, cpumult int, plain bool) (func(...int), []gbd.IntVar) {
	var cpu gbd.IntVar
	x := gbd.IntVar{VarName: "x", From: xfrom, To: xto, Multiple: xmultiple}
	if cpumult > 0 {
		cpu = gbd.IntVar{VarName: "routines", From: cpufrom, To: cputo, Multiple: cpumult}
	} else {
		cpu = gbd.IntVar{VarName: "routines", From: cpufrom, To: cputo, Increment: cpuinc}
	}
	ivs := []gbd.IntVar{x, cpu}
	if plain {
		return plainTestBenchWrapper, ivs
	}
	return superTestBenchWrapper, ivs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test(x int, checkSqrt bool) (int, int, int) {
	/*
			  a+b+c=x
			  we know a2+b2=c2 and hence can reduce a variable;
			  a+b+sqt(a2+b2)=x
		    as these are all natural numbers, sqrt(a2+b2) must be an int.

		    so we check sqrt(a2+b2) as int first.
	*/
	//Regardless of how small a gets, b<c so even at a=0, b+c=x, b<x/2 and hence
	// so will a
	if checkSqrt {
		return sqrtTest(x)
	}
	return plainTest(x)
}

func sqrtTest(x int) (int, int, int) {
	for a := 0; a < x/2; a++ {
		for b := a + 1; b < x/2; b++ {
			m := math.Sqrt(float64(a*a + b*b))
			im := int(m)
			if float64(int(m)) != m {
				continue
			}
			y := a + b + im
			if x == y {
				return a, b, im
			}
			if y > x {
				break
			}
		}
	}
	return 0, 0, 0
}
func plainTestBenchWrapper(x ...int) {
	plainTest(x[0])
}
func plainTest(x int) (int, int, int) {
	for a := 0; a < x/2; a++ {
		for b := a + 1; b < x/2; b++ {
			m := math.Sqrt(float64(a*a + b*b))
			im := int(m)

			y := a + b + im
			if x == y {
				//This will only run once if we are checking for sqrt, so no need to if
				// again
				if float64(im) != m {
					continue
				}
				return a, b, im
			}
			if y > x {
				break
			}
		}
	}
	return 0, 0, 0
}

func superTestBenchWrapper(x ...int) {
	_, _, _ = superTest(x[0], x[1])
}

func superTest(x, routines int) (int, int, int) {
	if routines == 0 {
		routines = 1
	}
	wg := sync.WaitGroup{}
	resultChan := make(chan int, 3)
	wg.Add(routines)
	quit := false
	z := 1
	for i := 0; i < routines; i++ {
		if i == routines-1 {
			go subtest(x, i*(x/routines), x, &wg, &quit, resultChan)
		} else {
			go subtest(x, i*(x/routines), (i+1)*(x/routines)-z, &wg, &quit, resultChan)
		}
	}
	wg.Wait()
	return <-resultChan, <-resultChan, <-resultChan
}

func subtest(x, amin, amax int, wg *sync.WaitGroup, quit *bool, result chan int) {
	//fmt.Println(x, amin, amax, *quit)
	for a := amin; a <= amax && !*quit; a++ {
		for b := a + 1; b < x/2 && !*quit; b++ {
			//fmt.Println(amin, amax, a, b, *quit)
			m := math.Sqrt(float64(a*a + b*b))
			im := int(m)
			y := a + b + im
			if x == y {
				if float64(im) != m {
					continue
				}
				*quit = true
				result <- a
				result <- b
				result <- im
				wg.Done()
				return
			}
			if y > x {
				break
			}
		}
	}
	wg.Done()
	return
}
