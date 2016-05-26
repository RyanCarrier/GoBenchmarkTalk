package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var tests int
var f = os.Stdin

func main() {
	fmt.Fscan(f, &tests)
	for i := 0; i < tests; i++ {
		fmt.Println(solveB(get(f)))
	}
}

func solveA(dolla int, items []int) (int, int) {
	var index int
	half := dolla / 2
	data := map[int]int{}
	for i := range items {
		if items[i] > half {
			index = dolla - items[i]
		} else {
			index = items[i]
		}
		if j, found := data[index]; found {
			if items[i]+items[j] == dolla {
				return i + 1, j + 1
			}
		}
		data[index] = i
	}
	return -1, -1
}

func solveB(dolla int, items []int) (int, int) {
	items2 := append([]int{}, items...)
	sort.Ints(items)
	i := 0
	j := len(items) - 1
	for i < j {
		temp := items[i] + items[j]
		switch {
		case temp == dolla:
			ii := 0
			for _, a := range items2 {
				if a == items[i] {
					break
				}
				ii++
			}
			jj := len(items2) - 1
			for jj >= 0 {
				if items2[jj] == items[j] {
					break
				}
				jj--
			}
			return ii + 1, jj + 1
		case temp > dolla:
			j--
		case temp < dolla:
			i++
		}
	}
	return -1, -1
}

func solveC(dolla int, items []int) (int, int) {
	//fmt.Println(dolla, items)
	for i, ii := range items {
		for j, jj := range items {
			if i == j {
				continue
			}
			if ii+jj == dolla {
				return i + 1, j + 1
			}
		}
	}
	return -1, -1
}

func get(r io.Reader) (int, []int) {
	var total int
	var entries int
	fmt.Fscan(r, &total)
	fmt.Fscan(r, &entries)
	nums := make([]int, entries)
	for i := 0; i < entries; i++ {
		fmt.Fscan(r, &nums[i])
	}
	return total, nums
}
