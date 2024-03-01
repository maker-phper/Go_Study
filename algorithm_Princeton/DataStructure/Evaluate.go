package main

// This code handles operations between integers, leaving many issues unresolved, including operator precedence

import (
	"fmt"
	"unicode"
)

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
	vals := Stack{}
	ops := Stack{}
	s := "(1 + ((21 - 1) * ((3 + 5) * (2 / 4))))"
	d := false
	for _, v := range s {

		if unicode.IsDigit(v) {
			val := float64(v - '0')
			if d {
				num := vals.Pop().(float64)
				val = num*10 + val
			}
			vals.Push(val)
			d = true
		} else {
			d = false
		}
		if v == '+' {
			ops.Push(v)
		} else if v == '*' {
			ops.Push(v)
		} else if v == '-' {
			ops.Push(v)
		} else if v == '/' {
			ops.Push(v)
		} else if v == ')' {
			op := ops.Pop()
			if op == '+' {
				vals.Push(vals.Pop().(float64) + vals.Pop().(float64))
			} else if op == '*' {
				vals.Push(vals.Pop().(float64) * vals.Pop().(float64))
			} else if op == '-' {
				vals.Push(-vals.Pop().(float64) + vals.Pop().(float64))
			} else if op == '/' {
				vals.Push(1 / vals.Pop().(float64) * vals.Pop().(float64))
			}
		}
	}
	fmt.Println(vals.Pop())
}
