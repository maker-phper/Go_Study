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
	if s.top.item == "" {
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
	as.n--
	return item
}

func (as *ArrStack) Push(item string) {
	as.n = as.n + 1
	as.items[as.n] = item
}

func main() {
	s := Stack{}
	s.Push("1")
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())

	size := 10
	as := &ArrStack{items: make([]string, size)}
	as.Push("a1")
	fmt.Println(as.Pop())
	fmt.Println(as.IsEmpty())
}
