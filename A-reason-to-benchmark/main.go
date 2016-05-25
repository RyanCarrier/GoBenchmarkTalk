package main

import "fmt"

//Set has a set of int which can be set and reset
type Set []int

func main() {
	set := Set{1, 2, 3, 4}
	fmt.Println("set:\n", set)
	set = set[0:3]
	fmt.Println("set after reslice:\n", set)
	//fmt.Println(set[3])
	clone := set[0:4]
	clone2 := set
	clone3 := clone2[0:4]
	//fmt.Println("set index 3\n", len(set), cap(set), set[3])
	fmt.Println("set index 3\n", len(set), cap(set))
	fmt.Println("clone index 3\n", len(clone), cap(clone), clone[3])
	//fmt.Println("clone2 index 3\n", len(clone2), cap(clone2), clone2[3])
	fmt.Println("clone2 index 3\n", len(clone2), cap(clone2))
	fmt.Println("clone3 index 3\n", len(clone3), cap(clone3), clone3[3])

}

//https://github.com/golang/go/wiki/SliceTricks
