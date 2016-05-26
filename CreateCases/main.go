package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

var filename = "test"
var base = 10

func main() {
	for x := 0; x < 7; x++ {
		create(10, x)
	}
	filename = "2test"
	for x := 0; x < 10; x++ {
		create(100000+x*100000, 0)
	}
}

func create(coins, n int) {

	tests := 1
	dolla := 1000
	half := (dolla - 2) / 2
	for x := 0; x < n; x++ {
		coins = coins * base
	}
	I, J := rand.Intn(coins-1), rand.Intn(coins-1)
	infile := filename + strconv.Itoa(coins) + ".in"
	os.Remove(infile)
	f, _ := os.OpenFile(infile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	f.WriteString(strconv.Itoa(tests) + "\n")
	f.WriteString(strconv.Itoa(dolla) + "\n")
	f.WriteString(strconv.Itoa(coins) + "\n")
	for i := 0; i < coins; i++ {
		if i == I || i == J {
			f.WriteString(strconv.Itoa(dolla/2) + " ")
		}
		f.WriteString(strconv.Itoa(rand.Intn(half)) + " ")
	}
	outfile := filename + strconv.Itoa(coins) + ".out"
	os.Remove(outfile)
	ioutil.WriteFile(outfile, []byte(strconv.Itoa(I)+" "+strconv.Itoa(J)), 0666)
}
