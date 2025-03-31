package main

import "fmt"

func main() {
	queue := &Queue[int]{}
	queue.Enqueue(4)
	queue.Enqueue(10)
	fmt.Println("First element in queue", queue.Dequeue())
	queue.Enqueue(287356)
	queue.Enqueue(6)
	fmt.Println("First element in queue", queue.Dequeue())
	fmt.Println("First element in queue", queue.Dequeue())
	fmt.Println("First element in queue", queue.Dequeue())
	fmt.Println("First element in queue", queue.Dequeue())
	queue.Enqueue(6)
	queue.Enqueue(2)
	queue.Enqueue(8)
	queue.Enqueue(1)
	fmt.Println("First element in queue", queue.Dequeue())
	fmt.Println("First element in queue", queue.Dequeue())
	fmt.Println("First element in queue", queue.Dequeue())
}

type Queue[T any] struct {
	s1 []T
	s2 []T
}

func (q *Queue[T]) Enqueue(element T) {
	q.s1 = append(q.s1, element)
}

func (q *Queue[T]) Dequeue() T {
	var zero T
	if len(q.s2) == 0 && len(q.s1) == 0 {
		return zero
	} else if len(q.s2) == 0 {
		q.dataTransfer()
	}

	element := q.s2[0]
	q.s2 = q.s2[1:]
	return element
}

func (q *Queue[T]) dataTransfer() {
	q.s2 = append(q.s2, q.s1...)
	q.s1 = nil
}
