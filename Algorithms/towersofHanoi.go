package main

import "fmt"

func towersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	tohUtil(num, "A", "C", "B")
}

func tohUtil(n int, from string, to string, temp string) {
	if n < 1 {
		return
	}
	tohUtil(n-1, from, temp, to)
	fmt.Println("Move disk", n, "from peg", from, "to peg", to)
	tohUtil(n-1, temp, to, from)

}

func main() {
	towersOfHanoi(3)
}
