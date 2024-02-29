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

func (q *Queue) Enqueue(item string) {
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

func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}
	item := q.first.item
	q.first = q.first.next
	return item
}

type ArrQueue struct {
	head  int
	tail  int
	items []string
}

func NewArrQueue() *ArrQueue {
	return &ArrQueue{
		items: make([]string, 2),
		tail:  -1,
	}
}

func (aq *ArrQueue) IsEmpty() bool {
	if aq.tail < aq.head {
		return true
	}
	return false
}

func (aq *ArrQueue) Enqueue(str string) {
	tail := aq.tail + 1
	if tail == len(aq.items) {
		aq.Resize(len(aq.items) * 2)
	}
	aq.tail = aq.tail + 1
	aq.items[aq.tail] = str
}

func (aq *ArrQueue) Dequeue() string {
	if aq.IsEmpty() {
		return ""
	}
	item := aq.items[aq.head]
	aq.items[aq.head] = ""
	aq.head++
	if (aq.tail-aq.head)*4 == cap(aq.items) {
		aq.Resize(cap(aq.items) / 2)
	}
	return item
}

func (aq *ArrQueue) Resize(size int) {
	newItems := make([]string, size)
	copy(newItems, aq.items[aq.head:aq.tail+1])
	aq.tail = aq.tail - aq.head
	aq.head = 0
	aq.items = newItems
}

func main() {
	newAq := NewArrQueue()
	fmt.Println(newAq.IsEmpty())
	newAq.Enqueue("a1")
	newAq.Enqueue("a2")
	fmt.Println(newAq.Dequeue())
	fmt.Println(newAq.Dequeue())
	fmt.Println(newAq.IsEmpty())
	
	nq := Queue{}
	fmt.Println(nq.IsEmpty())
	nq.Enqueue("a1")
	nq.Enqueue("a2")
	fmt.Println(nq.Dequeue())
	fmt.Println(nq.Dequeue())
	fmt.Println(nq.IsEmpty())
}
