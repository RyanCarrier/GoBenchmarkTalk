package main

import (
	"fmt"
	"strconv"
	"strings"
)

var tests int

func main() {
	fmt.Scan(&tests)
	for i := 0; i < tests; i++ {
		get()
	}
}

func solve(int, []int) (int, int) {

	return 0, 0
}

func get() (int, []int) {
	var total int
	var entries int
	var line string
	fmt.Scan(&total)
	fmt.Scan(&entries)
	fmt.Scan(&line)
	numstrings := strings.Fields(line)
	nums := make([]int, entries)
	for i, s := range numstrings {
		nums[i], _ = strconv.Atoi(s)
	}
	return total, nums
}
