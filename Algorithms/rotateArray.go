package main

import "fmt"

//Given a list, you need to rotate its elements K number of times.
//For example, a list [10,20,30,40,50,60] rotate by 2 positions to [30,40,50,60,10,20]

func rotateArray(mas []int, k int) {
	len := len(mas)
	reverseArray(mas, 0, k-1)
	reverseArray(mas, k, len-1)
	reverseArray(mas, 0, len-1)

}

func reverseArray(mas []int, start int, stop int) {
	i := start
	j := stop

	for i < j {
		mas[i], mas[j] = mas[j], mas[i]
		i++
		j--
	}
}

func main() {
	mas := []int{10, 20, 30, 40, 50, 60}
	rotateArray(mas, 2)
	fmt.Println(mas)
}
