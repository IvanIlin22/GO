package main

import "fmt"

func maxSubArray(mas []int) int {
	maxSoFar := 0
	maxHereEnd := 0

	for _, val := range mas {
		maxHereEnd += val
		if maxHereEnd < 0 {
			maxHereEnd = 0
		}
		if maxSoFar < maxHereEnd {
			maxSoFar = maxHereEnd
		}
	}
	return maxSoFar
}

func main() {

	mas := []int{1, -2, 5, 4, -4, 6, -14, 8, 2}
	max := maxSubArray(mas)
	fmt.Println(max)
}
