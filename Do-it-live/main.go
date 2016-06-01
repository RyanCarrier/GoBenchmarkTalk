package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var tests int

func main() {
	f := os.Stdin
	fmt.Fscan(f, &tests)
	for test := 0; test < tests; test++ {
		fmt.Println(solve2C(get(f)))
	}
}

func solveA(dolla int, items []int) (int, int) {
	for i, item := range items {
		for j, item2 := range items {
			if i != j && item+item2 == dolla {
				return i + 1, j + 1
			}
		}
	}
	return -1, -1
}

func solve2A(dolla int, items []int) (int, int) {
	for i, item := range items {

		l := len(items)
		for j := i + 1; j < l; j++ {
			if i != j && item+items[j] == dolla {
				return i + 1, j + 1
			}
		}
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
		if temp == dolla {
			Ifinal := -1
			Jfinal := -1
			for I, item := range items2 {
				if Ifinal == -1 && items[i] == item {
					Ifinal = I
				}
				if items[j] == item {
					Jfinal = I
				}
			}
			return Ifinal + 1, Jfinal + 1
		}
		if temp < dolla {
			i++
		} else {
			j--
		}
	}
	return -1, -1
}

func solveC(dolla int, items []int) (int, int) {
	data := map[int]int{}
	half := dolla / 2
	var index int
	for i, item := range items {
		if item > half {
			index = dolla - item
		} else {
			index = item
		}
		if item2, ok := data[index]; ok {
			if items[item2]+item == dolla {
				return item2 + 1, i + 1
			}
		}
		data[index] = i
	}
	return -1, -1
}

func solve2C(dolla int, items []int) (int, int) {
	var index int
	data := make([]int, (dolla/2)+1)
	for i, it := range items {
		if it > dolla/2 {
			index = dolla - it
		} else {
			index = it
		}
		if data[index] == 0 {
			data[index] = i + 1
		} else {
			if items[data[index]-1]+items[i] == dolla {
				return data[index], i + 1
			}
		}
	}
	return -1, -1
}

func get(f io.Reader) (dolla int, items []int) {
	var itemsN int
	fmt.Fscan(f, &dolla, &itemsN)
	items = make([]int, itemsN)
	for i := range items {
		fmt.Fscan(f, &items[i])
	}
	return
}
