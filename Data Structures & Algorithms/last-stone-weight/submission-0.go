func lastStoneWeight(stones []int) int {
	h := &maxHeap{}
	heap.Init(h)
	for _, num := range stones {
		heap.Push(h, num)
	}
	
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x > y {
			heap.Push(h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() any {
	val := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return val
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
