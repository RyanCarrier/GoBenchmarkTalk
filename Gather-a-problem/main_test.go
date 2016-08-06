package main

import t "testing"

func benchFib(b *t.B, n int) {
	for i := 0; i < b.N; i++ {
		fib(int64(n))
	}
}

func benchPrime(b *t.B, n int) {
	for i := 0; i < b.N; i++ {
		prime(n)
	}
}
func bench(b *t.B, p int) {
	for i := 0; i < b.N; i++ {
		switch p {
		case 1:
			p1()
		case 2:
			p2()
		case 3:
			p3()
		case 4:
			p4()
		case 6:
			p6()
		case 7:
			p7()
		case 10:
			p10()
		}
	}
}
func BenchmarkAll(b *t.B) {
	for i := 0; i < b.N; i++ {
		all()
	}
}
func BenchmarkConcurrentAll(b *t.B) {
	for i := 0; i < b.N; i++ {
		concurrentAll()
	}
}
func BenchmarkP1(b *t.B)  { bench(b, 1) }
func BenchmarkP2(b *t.B)  { bench(b, 2) }
func BenchmarkP3(b *t.B)  { bench(b, 3) }
func BenchmarkP4(b *t.B)  { bench(b, 4) }
func BenchmarkP6(b *t.B)  { bench(b, 6) }
func BenchmarkP7(b *t.B)  { bench(b, 7) }
func BenchmarkP10(b *t.B) { bench(b, 10) }

func BenchmarkFib1(b *t.B)  { benchFib(b, 1) }
func BenchmarkFib10(b *t.B) { benchFib(b, 10) }
func BenchmarkFib20(b *t.B) { benchFib(b, 20) }
func BenchmarkFib25(b *t.B) { benchFib(b, 25) }
func BenchmarkFib30(b *t.B) { benchFib(b, 30) }

func BenchmarkPrime1(b *t.B)    { benchPrime(b, 1) }
func BenchmarkPrime10(b *t.B)   { benchPrime(b, 10) }
func BenchmarkPrime50(b *t.B)   { benchPrime(b, 50) }
func BenchmarkPrime100(b *t.B)  { benchPrime(b, 100) }
func BenchmarkPrime1000(b *t.B) { benchPrime(b, 1000) }

func TestFib(t *t.T) {
	indexes := []int64{0, 1, 10, 45}
	answers := []int64{0, 1, 55, 1134903170}
	for i := range indexes {
		if answers[i] != fib(indexes[i]) {
			t.Error("fib", indexes[i], "failed\nGOT:", fib(indexes[i]), "\nWANT:", answers[i])
		}
	}
}

func TestPrime(t *t.T) {
	indexes := []int{1, 10, 100} //, 1000}
	answers := []int{2, 29, 541} //, 7919}
	for i := range indexes {
		if answers[i] != prime(indexes[i]) {
			t.Error("prime", indexes[i], "failed\nGOT:", prime(indexes[i]), "\nWANT:", answers[i])
		}
	}
}
