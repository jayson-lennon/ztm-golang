package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]
		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}
	var empty V
	return empty, false
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{items: make(map[P][]V), priorities: priorities}
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
}
