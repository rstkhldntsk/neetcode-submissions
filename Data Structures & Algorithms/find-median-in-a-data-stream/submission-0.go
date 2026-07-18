type MedianFinder struct {
	lo *MaxHeap
	hi *MinHeap
}

func Constructor() MedianFinder {
	lo := &MaxHeap{}
	hi := &MinHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{lo, hi}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, num)
	heap.Push(this.hi, heap.Pop(this.lo))
	if this.hi.Len() > this.lo.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64((*this.lo)[0])
	}
	return float64((*this.lo)[0] + (*this.hi)[0]) / 2.
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
