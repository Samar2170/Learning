package main

type Queue struct {
	Items   []int
	MaxSize int
}

func (q *Queue) Enqueue(val int) {
	if len(q.Items) == q.MaxSize {
		panic("Max Size reached")
	}
	q.Items = append(q.Items, val)
}

func (q *Queue) Dequeue() int {
	popped := q.Items[0]
	q.Items = q.Items[1:]
	return popped
}

func (q *Queue) Peek() int {
	return q.Items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.Items) == q.MaxSize
}

func main() {
	q := Queue{MaxSize: 5}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
