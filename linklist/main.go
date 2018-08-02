package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func (head *Node) push(v int) {
	newNode := newNode(v)
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func newNode(v int) *Node {
	newNode := new(Node)
	newNode.value = v
	return newNode
}

func (head *Node) String() string {
	temp := head
	s := ""
	for temp != nil {
		s += fmt.Sprintf(" {%d} ", temp.value)
		temp = temp.next
	}
	return s
}

func main() {
	head := newNode(1)
	slice := []int{2, 3, 4, 5, 6, 7, 8}
	for _, elem := range slice {
		head.push(elem)
	}

	fmt.Println(head)
}
