package main

import "fmt"

func gcd(m int, n int) int {
	if m < n {
		return gcd(n, m)
	}
	if m%n == 0 {
		return n
	}
	return gcd(n, m%n)
}

func main() {
	fmt.Println(gcd(24, 17))
}