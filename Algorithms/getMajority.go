package main

import (
	"fmt"
	"sort"
)

func getMagority(mas []int) (int, bool) {
	var max int = 0
	var maxCount int = 1
	var count int = 1
	for i := 0; i < len(mas); i++ {
		count = 1
		for j := i + 1; j < len(mas); j++ {
			if mas[i] == mas[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = mas[i]
		}
	}
	if maxCount > len(mas)/2 {
		return max, true
	}
	return 0, false
}

func getMagority2(mas []int) (int, bool) {
	var count int = 0

	sort.Ints(mas)

	for i := 0; i < len(mas); i++ {
		if mas[i] == mas[len(mas)/2] {
			count++
		}
	}
	if count > len(mas)/2 {
		return mas[len(mas)/2], true

	}
	return 0, false
}

func getMagority3(mas []int) (int, bool) {
	var count int = 1
	index := 0

	for i := 1; i < len(mas); i++ {
		if mas[i] == mas[index] {
			count++
		} else {
			count--
		}
		if count == 0 {
			index = i
			count = 1
		}
	}
	count = 0
	for i := 0; i < len(mas); i++ {
		if mas[index] == mas[i] {
			count++
		}
	}
	if count > len(mas)/2 {
		return mas[index], true
	}
	return 0, false
}

func main() {
	mas := []int{5, 5, 3, 1, 5, 5, 1, 5, 3, 5, 6}
	fmt.Println(getMagority(mas))
	fmt.Println(getMagority2(mas))
	fmt.Println(getMagority3(mas))
}