package main

import (
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var filename = "test"

func main() {
	for x := 0; x < 7; x++ {
		create(10, x, true)
	}
}

func create(coinBase, n int, cheat bool) {
	tests := 1

	//halfOdd := half + 3
	coins := int64(coinBase)
	for x := 0; x < n; x++ {
		coins = coins * int64(coinBase)
	}
	var dolla int64
	if 2*coins >= math.MaxInt32 {
		dolla = 2 * ((int64(math.MaxInt32) / 2) - 1)
	} else {
		dolla = 2 * coins
	}
	half := ((dolla - 2) / 2)
	var I, J int64
	if cheat {
		I, J = coins-2, coins-1
	} else {
		I, J = rand.Int63n(coins-2), rand.Int63n(coins-2)
		if I == J {
			J++
		}
	}
	infile := filename + strconv.FormatInt(coins, 10) + ".in"
	os.Remove(infile)
	f, _ := os.OpenFile(infile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	f.WriteString(strconv.Itoa(tests) + "\n")
	f.WriteString(strconv.FormatInt(dolla, 10) + "\n")
	f.WriteString(strconv.FormatInt(coins, 10) + "\n")
	for i := int64(0); i < coins; i++ {
		if i == I || i == J {
			f.WriteString(strconv.FormatInt(dolla/2, 10) + " ")
		} else {
			if i%2 == 0 {
				f.WriteString(strconv.FormatInt(rand.Int63n(half/2)*2, 10) + " ")
			} else {
				f.WriteString(strconv.FormatInt((rand.Int63n(dolla/4)*2)+(2*(dolla/4))+1, 10) + " ")
			}
		}
	}
	outfile := filename + strconv.FormatInt(coins, 10) + ".out"
	os.Remove(outfile)
	ioutil.WriteFile(outfile, []byte(strconv.FormatInt(I+1, 10)+" "+strconv.FormatInt(J+1, 10)), 0666)
}
