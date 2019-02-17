package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(n int) {
	node := &Node{value: n}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
}

func (n Node) Next() *Node {
	return n.next
}

func main() {
	l := List{}
	l.Push(10)
	l.Push(20)
	l.Push(30)

	n := l.First()

	for {
		fmt.Println(n.value)

		n = n.Next()

		if n == nil {
			break
		}
	}
}
