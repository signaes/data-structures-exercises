package main

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	size := len(s.items)

	if size == 0 {
		return -1
	}

	item, items := s.items[size-1], s.items[0:size-1]
	s.items = items

	return item
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
