package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walked(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walked(t.Left, ch)
	ch <- t.Value
	Walked(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	Walked(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {
			return false
		}
	}
	return true
}

/*func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var1, ok1 := <-ch1
	var2, ok2 := <-ch2
	for {
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if var1 != var2 {
			return false
		} else {
			return true
		}
	}
	return true
}*/

/*func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		if <- ch1 != <- ch2 {
			return false
		} else {
			return true
		}
	}
	return true
}*/

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}