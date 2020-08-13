package main

import (
	"fmt"
)

func binary(data []int, val int) int {
	min := 0
	var max int = len(data) - 1

	for min <= max {
		var mid int = (min + max) / 2
		if data[mid] == val {
			return val
		} else if data[mid] < val {
			min = mid + 1
		} else if data[mid] > val {
			max = mid - 1
		}
	}
	return 0
}

func binaryRecursive(data []int, min int, max int, val int) int {
	if min > max {
		return 0
	}

	var mid int = (min + max) / 2

	if data[mid] == val {
		return val
	} else if data[mid] > val {
		return binaryRecursive(data, min, mid-1, val)
	} else {
		return binaryRecursive(data, mid+1, max, val)
	}
}

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(binary(data, 0))
	fmt.Println(binaryRecursive(data, 0, len(data), 4))
}
