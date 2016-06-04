package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var tests int
var f = os.Stdin

/*
	input format;
#tests
#dollaValue
#numberOfItems
#item1 #item2 .... #item[numberOfItems]

eg:
2
10
3
1 5 5
4
4
1 2 3 5
*/

func main() {
	//Later we will be reading from file not stdin, so we use Fscan instead of Scan;
	fmt.Fscan(f, &tests)
	for i := 0; i < tests; i++ {
		fmt.Println(solveB(get(f)))
	}
}

//O(N) method
//we loop through each item, storing it in a map, checking for collisions first
//index by the difference from the ends
func solveA(dolla int, items []int) (int, int) {
	var index int
	half := dolla / 2
	data := map[int]int{}
	for i := range items {
		//we want differences from each end to colide;
		//So for 100, we want 99 and 1 to both index 1
		if items[i] > half {
			index = dolla - items[i]
		} else {
			index = items[i]
		}
		//Check for colisions
		if j, found := data[index]; found {
			//Make sure the collision is the one we want, don't want to return
			//1 and 1, rather than 1 and 99
			if items[i]+items[j] == dolla {
				return i + 1, j + 1
			}
		}
		data[index] = i
	}
	return -1, -1
}

//O(N)
//Only difference from A, is use of array instead of map.
//Indexed from 1, so we know if the entry is set or not, map will tell us
//if there is something there, array won't.
func solveA2(dolla int, items []int) (int, int) {
	var index int
	data := make([]int, (dolla/2)+1)
	half := dolla / 2
	for i, it := range items {
		if it > half {
			index = dolla - it
		} else {
			index = it
		}
		if data[index] == 0 {
			data[index] = i + 1
		} else {
			//-1 from the index as we are (storeing) indexing from 1 (instead of 0)
			if items[data[index]-1]+items[i] == dolla {
				return data[index], i + 1
			}
		}
	}
	return -1, -1
}

//O(Nlog(N)), could be O(N^2) at worst case.
//This one is weird, the results appear O(N), until we hit 10M
//It's important to note that even though we iterate through the (potentially)
// entire list, this is only done once, and hence quicksort+iteration will equal
// O(Nlog(N))+O(N) and as the 'heavier' one will take dominance, it results in
// O(Nlog(N)).
//This method sorts the input, then checks from the extremities, moving in until
// it finds a solution. This will skirt around the solution until it finds it,
// meaning, it will start by joining the smallest and largest number(s). If too
// big, then take the next largest number, and opposite for smallest.
//Even though we are moving 2 indicies, together they will only move N times
// maximum.
func solveB(dolla int, items []int) (int, int) {
	//Create a backup of the original order
	items2 := append([]int{}, items...)
	//Don't need to create special type to sort
	sort.Ints(items)
	i := 0
	j := len(items) - 1
	//While we aren't at the middle
	for i < j {
		//Cheaper to store this than recalculate consttantly.
		temp := items[i] + items[j]
		switch {
		case temp == dolla:
			Ifinal := -1
			Jfinal := -1
			for I, item := range items2 {
				//find original i, aslong as it isn't set yet
				if Ifinal == -1 && items[i] == item {
					Ifinal = I
				}
				//find original j
				if items[j] == item {
					Jfinal = I
				}
			}
			return Ifinal + 1, Jfinal + 1
		case temp > dolla:
			//If we are too big, take a smaller biggest number.
			j--
		case temp < dolla:
			//If we are too small, take a larger smallest number.
			i++
		}
	}
	return -1, -1
}

//O(N^2)
//C is the 'easy' way of solving, very innefficient, but doesn't use any extra
// memory to solve.
//For each item, check if you hit dolla when adding it to every other item.
func solveC(dolla int, items []int) (int, int) {
	for i, ii := range items {
		for j, jj := range items {
			//Don't compare if they are the same
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

//O(N^2)
//C2 is the same as C1, but ignores already compared items.
//This will mean the second loop, on average (in the worst case scenario)
// is N/2, which doesn't make a great change, still O(N), but faster.
func solveC2(dolla int, items []int) (int, int) {
	var j int
	l := len(items)
	for i, ii := range items {
		for j = i + 1; j < l; j++ {
			if ii+items[j] == dolla {
				return i + 1, j + 1
			}
		}
	}
	return -1, -1
}

//get the case from the input
//reader used so we can parse stdin or a file (from os.Open())
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
