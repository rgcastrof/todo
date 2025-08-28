package tasks

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(id int) {
	q.items = append(q.items, id)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
