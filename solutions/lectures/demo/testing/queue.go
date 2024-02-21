package queue

type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}
	q.items = append(q.items, item)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.items) > 0 {
		next := q.items[0]
		q.items = q.items[1:]
		return next, true
	} else {
		return 0, false
	}
}
