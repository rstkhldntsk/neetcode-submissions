type MyCircularQueue struct {
    size int
	cap int
	sentinel *Node
}

type Node struct {
	v int
	prev, next *Node
}

func Constructor(k int) MyCircularQueue {
	sentinel := &Node{}
	sentinel.next, sentinel.prev = sentinel, sentinel
    return MyCircularQueue{
		size: 0,
		cap: k,
		sentinel: sentinel,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.size == q.cap {
		return false
	}
	n := &Node{v: value, prev: q.sentinel.prev, next: q.sentinel}
	q.sentinel.prev.next, q.sentinel.prev = n, n
	q.size++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
    if q.size == 0 {
		return false
	}
	q.sentinel.next = q.sentinel.next.next
	q.sentinel.next.prev = q.sentinel
	q.size--
	return true
}

func (q *MyCircularQueue) Front() int {
    if q.size > 0 {
		return q.sentinel.next.v
	}
	return -1
}


func (q *MyCircularQueue) Rear() int {
	if q.size > 0 {
		return q.sentinel.prev.v
	}
	return -1
}

func (q *MyCircularQueue) IsEmpty() bool {
    return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
    return q.size == q.cap
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 