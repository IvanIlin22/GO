package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head  *Node
	count int
}

func (list *List) addHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *List) addTail(value int) {
	new := &Node{value, nil}
	cur := list.head

	if cur == nil {
		list.head = new
		list.count++
		return
	}

	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new
	list.count++
}

func (list *List) Print() {
	cur := list.head
	for cur != nil {
		fmt.Print(cur.val, " ")
		cur = cur.next
	}
	fmt.Println("")
}

func (list *List) sortedInsert(val int) {
	new := &Node{val, nil}
	cur := list.head

	if cur == nil || cur.val > val {
		new.next = list.head
		list.head = new
		list.count++
		return
	}
	for cur.next != nil && cur.next.val < val {
		cur = cur.next
	}
	new.next = cur.next
	cur.next = new
	list.count++
}

func (list *List) isPresent(value int) bool {
	cur := list.head
	for cur != nil {
		if cur.val == value {
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *List) removeHead() (int, bool) {
	cur := list.head

	if cur == nil {
		return 0, false
	}
	val := cur.val
	list.head = cur.next
	list.count--
	return val, true
}

func (list *List) deleteNode(value int) bool {
	cur := list.head

	if cur == nil {
		return false
	}
	if cur.val == value {
		list.head = cur.next
		list.count--
		return true
	}
	for cur.next != nil {
		if cur.next.val == value {
			cur.next = cur.next.next
			list.count--
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *List) deleteNodes(value int) {
	cur := list.head

	for cur != nil && cur.val == value {
		list.head = cur.next
		cur = list.head
		list.count--
	}

	for cur.next != nil {
		nextNode := cur.next
		if nextNode.val == value && nextNode != nil {
			cur.next = nextNode.next
			list.count--
		} else {
			cur = cur.next
		}
	}
}

func main() {
	list := &List{}

	for i := 0; i < 4; i++ {
		list.addTail(i)
	}

	list.sortedInsert(3)
	list.sortedInsert(8)
	list.addHead(100)
	ok := list.deleteNode(8)
	fmt.Println("ok = ", ok)
	list.deleteNodes(3)
	list.Print()
	fmt.Println(list.count)

}
