func longestDiverseString(a int, b int, c int) string {
	result := make([]byte, 0)
	h := &MaxHeap{}
	heap.Init(h)
	heap.Push(h, Pair{char: 'a', count: a})
	heap.Push(h, Pair{char: 'b', count: b})
	heap.Push(h, Pair{char: 'c', count: c})
	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		if len(result) > 1 && result[len(result)-1] == p.char && result[len(result)-2] == p.char {
			if h.Len() == 0 {
				break
			}
			p2 := heap.Pop(h).(Pair)
			if p2.count > 0 {
				result = append(result, p2.char)
				p2.count--
				heap.Push(h, p2)
			}
			heap.Push(h, p)
		} else {
			if p.count > 0 {
				result = append(result, p.char)
				p.count--
				heap.Push(h, p)
			}
		}
	}
	return string(result)
}

type Pair struct {
	char  byte
	count int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count > h[j].count
	}
	return h[i].char < h[j].char
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}
