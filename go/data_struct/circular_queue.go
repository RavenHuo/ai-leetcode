package data_struct

type CircularQueue struct {
	size  int
	cap   int
	head  int
	tail  int
	queue []int
}

func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		size:  0,
		cap:   cap,
		head:  0,
		tail:  0,
		queue: make([]int, cap),
	}
}

func (q *CircularQueue) EnQueue(n int) bool {
	if q.size == q.cap {
		return false
	}
	q.queue[q.tail] = n
	q.tail = (q.tail + 1) % q.cap
	q.size = q.size + 1
	return true
}

// 出队
func (q *CircularQueue) DoQueue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	v := q.queue[q.head]
	q.head = (q.head + 1) % q.cap
	q.size = q.size - 1
	return v, true
}
