package main

import "fmt"

func main() {
	set := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(set, len(set), cap(set))
	set = set[0:5]
	set := append([]int{}, set[0:5]...)
  var set2 []int{}
  copy(set2,set[0:5])
	fmt.Println(set, len(set), cap(set))
	//fmt.Println(set[5])
	clone := set[0:6]
	fmt.Println(clone[5])
}
