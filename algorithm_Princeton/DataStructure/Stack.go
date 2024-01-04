package main

import "fmt"

type Node struct {
	item string
	next *Node
}

type Stack struct {
	top *Node
}

type ArrStack struct {
	items []string
	n     int
}

func NewArrStack() *ArrStack {
	return &ArrStack{
		items: make([]string, 2),
		n:     -1,
	}
}

func (s *Stack) Push(item string) {
	newNode := &Node{item: item, next: s.top}
	s.top = newNode
}

func (s *Stack) Pop() string {
	if s.top == nil {
		return ""
	}
	item := s.top.item
	s.top = s.top.next
	return item
}

func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (as *ArrStack) IsEmpty() bool {
	if as.n == 0 {
		return true
	}
	return false
}

func (as *ArrStack) Pop() string {
	item := as.items[as.n]
	as.items[as.n] = ""
	as.n--
	if as.n+1 == len(as.items)/4 {
		as.Resize(len(as.items) / 2)
	}
	return item
}

func (as *ArrStack) Push(item string) {
	if as.n+1 == len(as.items) {
		as.Resize(2 * len(as.items))
	}
	as.n = as.n + 1
	as.items[as.n] = item

}

func (as *ArrStack) Resize(size int) {
	newItems := make([]string, size)
	copy(newItems, as.items)
	as.items = newItems
}

func main() {
	s := Stack{}
	s.Push("1")
	s.Push("2")
	s.Push("3")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())

	as := NewArrStack()
	as.Push("a1")
	as.Push("a2")
	as.Push("a3")
	as.Push("a4")
	fmt.Println(as)
	fmt.Println(as.Pop())
	fmt.Println(as.Pop())
	fmt.Println(as.Pop())
	fmt.Println(len(as.items))
	fmt.Println(as.IsEmpty())
}
