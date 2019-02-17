package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(n int) {
	node := &Node{value: n}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Find(v int) *Node {
	var found *Node
	n := l.First()

	for {
		if n.value == v {
			found = n
			break
		}

		n = n.Next()

		if n == nil {
			break
		}
	}

	return found
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Prev() *Node {
	return n.prev
}

func main() {
	l := List{}
	l.Push(10)
	l.Push(20)
	l.Push(30)

	fmt.Println(">>")

	n := l.First()

	for {
		fmt.Println(n.value)

		n = n.Next()

		if n == nil {
			break
		}
	}

	fmt.Println("<<")

	n = l.Last()

	for {
		fmt.Println(n.value)

		n = n.Prev()

		if n == nil {
			break
		}
	}

	fmt.Println("#Find")
	fmt.Println(l.Find(20))
	fmt.Println(l.Find(20).value)
}
