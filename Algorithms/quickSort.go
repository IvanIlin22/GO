package main

import (
	"fmt"
)

func swap(list []int, first int, sec int) {
	list[first], list[sec] = list[sec], list[first]
}

func more(x int, y int) bool {
	if x > y {
		return true
	}
	return false
}

func quickSort(list []int, comp func(int, int) bool) {
	size := len(list)
	quickSortUtil(list, 0, size-1, comp)
}

func quickSortUtil(list []int, lower int, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	start := lower
	stop := upper
	pivot := list[lower]

	for lower < upper {
		for comp(list[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(list[upper], pivot) && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(list, upper, lower)
		}
	}
	swap(list, upper, start)
	// upper is the pivot position
	quickSortUtil(list, start, upper-1, comp) // pivot -1 is the upper for left sub array.
	quickSortUtil(list, upper+1, stop, comp)  // pivot + 1 is the lower for right sub array.
}


func main() {
	//data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	data := []int{6, 4, 5, 1, 8, 2, 9, 7, 10, 3}
	//data := []int{10, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	quickSort(data, more)
	fmt.Println(data)
}
