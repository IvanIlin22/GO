package main

import (
	"fmt"
	"sort"
)

//The Time Complexity is O(n^2) and Space Complexity is O(1)
func getMaxRepeat(data []int) int {
	count := 1
	maxCount := 1
	max := data[0]
	size := len(data)

	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = data[i]
		}
	}
	return max
}

//The Time Complexity is O(nlog(n)) and Space Complexity is O(1)
func getMaxRepeat1(data []int) int {
	sort.Ints(data)

	count := 1
	countMax := 1
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] {
			count++
		} else {
			count = 1
		}
		if count > countMax {
			countMax = count
			max = data[i]
		}
	}
	return max
}

//The Time Complexity is O(n) and Space Complexity is O(n)
func getMaxRepeat2(data []int, dataRange int) int {
	count := make([]int, dataRange)
	countMax := 1
	max := data[0]

	for i := 0; i < len(data); i++ {
		count[data[i]]++
		if count[data[i]] > countMax {
			countMax = count[data[i]]
			max = data[i]
		}
	}
	return max
}

func main() {
	data := []int{2, 1000, 2, 3, 4, 5, 2, 7, 8, 1000, 1000}

	fmt.Println(getMaxRepeat(data))
	fmt.Println(getMaxRepeat1(data))
	fmt.Println(getMaxRepeat2(data, 1001))
}
