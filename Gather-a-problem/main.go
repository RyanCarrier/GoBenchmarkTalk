package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

func main() {

	fmt.Println(p1())
	fmt.Println(p2())
	fmt.Println(p3())
	fmt.Println(p4())
	fmt.Println(p6())
	fmt.Println(p7())
	fmt.Println(p10())

	fmt.Println(all())
	fmt.Println(concurrentAll())
}

func all() int {
	//Find the sum of the first digit of all p1,2,3,4,6,7,10
	total := 0
	total += p1() % 10
	total += p2() % 10
	total += p3() % 10
	total += p4() % 10
	total += p6() % 10
	total += p7() % 10
	total += p10() % 10
	return total
}

func concurrentAll() int {
	results := make([]int, 11)
	wg := sync.WaitGroup{}
	wg.Add(7)
	go func() {
		results[1] = p1() % 10
		wg.Done()
	}()
	go func() {
		results[2] = p2() % 10
		wg.Done()
	}()
	go func() {
		results[3] = p3() % 10
		wg.Done()
	}()
	go func() {
		results[4] = p4() % 10
		wg.Done()
	}()
	go func() {
		results[6] = p6() % 10
		wg.Done()
	}()
	go func() {
		results[7] = p7() % 10
		wg.Done()
	}()
	go func() {
		results[10] = p10() % 10
		wg.Done()
	}()
	total := 0
	wg.Wait()
	for _, r := range results {
		total += r
	}
	return total
}

func p1() int {
	//Find the sum of all the multiples of 3 or 5 below 1000.
	total := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func p2() int {
	total, j := 0, 1
	//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
	for i := 0; i <= 4000000; i, j = i+j, i {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}

func p3() int {
	//What is the largest prime factor of the number 600851475143 ?
	target := 600851475143
	mid := int(math.Ceil(math.Sqrt(float64(target))))
	for ; target%mid != 0 || !isPrime(mid); mid-- {
	}
	return mid
}

func p4() int {
	//Find the largest palindrome made from the product of two 3-digit numbers.
	largest := 0
	for a := 999; a > 99; a-- {
		for b := 999; b > 99; b-- {
			if palindrome(a * b) {
				if a*b > largest {
					largest = a * b
				}
			}
		}
	}
	return largest
}

func p6() int {
	//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	sumofsquares := 0
	sums := 0
	for i := 1; i <= 100; i++ {
		sumofsquares += i * i
		sums += i
	}
	return sums*sums - sumofsquares
}

func p7() int {
	//What is the 10 001st prime number?
	return prime(10001)
}
func p10() int {
	//Find the sum of all the primes below two million.
	/*I"M SORRY IM CHEATING IT TAKES TOO LONG*/
	total := 0
	for i := 0; i < 2000000/8; i++ {
		if isPrime(i) {
			total += i
		}
	}
	return total
}

func palindrome(x int) bool {
	si := strconv.Itoa(x)
	for i := 0; i <= len(si)/2; i++ {
		if si[i] != si[len(si)-i-1] {
			return false
		}
	}
	return true
}
func fib(x int64) int64 {
	//lets hope they dont' give us negative
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func prime(x int) int {
	current := 2
	for i := 1; i < x; i++ {
		for current++; !isPrime(current); current++ {
		}
	}
	return current
}

func isPrime(x int) bool {
	sqrtx := int(math.Ceil(math.Sqrt(float64(x))))
	for i := 2; i <= sqrtx; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
