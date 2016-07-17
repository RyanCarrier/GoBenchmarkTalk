package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
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
	fmt.Println(runtime.NumCPU())
	fmt.Println(plainTest(10000))
	fmt.Println(superTest(10000, 10, true))
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

func superTest(x, routines int, quitEnabled bool) (int, int, int) {
	wg := sync.WaitGroup{}
	resultChan := make(chan int, 3)
	wg.Add(routines)
	if x%routines != 0 {
		//fmt.Print("Fk u")
		return 0, 0, 0
	}
	quit := false
	z := 1
	for i := 0; i < routines; i++ {
		if i == routines-1 {
			z = 0
		}
		go subtest(x, i*(x/routines), (i+1)*(x/routines)-z, &wg, quitEnabled, &quit, resultChan)
	}
	wg.Wait()
	return <-resultChan, <-resultChan, <-resultChan
}

func subtest(x, amin, amax int, wg *sync.WaitGroup, quitEnabled bool, quit *bool, result chan int) {
	//fmt.Println(x, amin, amax, *quit)
	for a := amin; a <= amax; a++ {
		if quitEnabled && *quit {
			//fmt.Println("QuitingEarly")
			break
		}
		for b := a + 1; b < x/2; b++ {
			if quitEnabled && *quit {
				break
			}
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
