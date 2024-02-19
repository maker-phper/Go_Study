package main

import "fmt"

type QNode struct {
	next *QNode
	item string
}

type Queue struct {
	first *QNode
	rear  *QNode
}

func (q *Queue) IsEmpty() bool {
	if q.first != nil {
		return false
	}
	return true
}

func (q *Queue) Push(item string) {
	newNode := &QNode{
		item: item,
	}
	if q.IsEmpty() {
		q.first = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

}

func (q *Queue) Pop() string {
	if q.IsEmpty() {
		return ""
	}
	item := q.first.item
	q.first = q.first.next
	return item
}

func main() {
	nq := Queue{}
	fmt.Println(nq.IsEmpty())
	nq.Push("a1")
	nq.Push("a2")
	fmt.Println(nq.Pop())
	fmt.Println(nq.Pop())
	fmt.Println(nq.IsEmpty())
}
