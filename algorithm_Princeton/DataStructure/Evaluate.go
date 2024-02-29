package main

import "fmt"

type Node struct {
	item interface{}
	next *Node
}

type Stack struct {
	top *Node
}

func (stack *Stack) IsEmpty() bool {
	if stack.top == nil {
		return true
	}
	return false
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	item := stack.top.item
	stack.top = stack.top.next
	return item
}

func (stack *Stack) Push(item interface{}) {
	newStack := &Node{
		item: item,
		next: stack.top,
	}
	stack.top = newStack
}

func main() {
	s := Stack{}
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push("b")
	s.Push("*")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
