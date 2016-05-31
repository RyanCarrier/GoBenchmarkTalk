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
		create(10, x, true)
	}
}

func create(coins, n int, cheat bool) {
	tests := 1
	dolla := 2 * coins
	half := ((dolla - 2) / 2)
	halfOdd := half + 1
	for x := 0; x < n; x++ {
		coins = coins * base
	}
	var I, J int
	if cheat {
		I, J = coins-2, coins-1
	} else {
		I, J = rand.Intn(coins-2), rand.Intn(coins-2)
		if I == J {
			J++
		}
	}
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
		} else {
			if i%2 == 0 {
				f.WriteString(strconv.Itoa(rand.Intn(half/2)*2) + " ")
			} else {
				f.WriteString(strconv.Itoa(rand.Intn(half/2)*2+halfOdd) + " ")
			}
		}
	}
	outfile := filename + strconv.Itoa(coins) + ".out"
	os.Remove(outfile)
	ioutil.WriteFile(outfile, []byte(strconv.Itoa(I+1)+" "+strconv.Itoa(J+1)), 0666)
}
