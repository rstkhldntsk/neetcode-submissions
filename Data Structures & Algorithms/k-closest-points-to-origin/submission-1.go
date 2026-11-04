func kClosest(points [][]int, k int) [][]int {
    h := make(minHeap, len(points))

    for i, p := range points {
        h[i] = P{
            point: p,
            dist:  dist(p),
        }
    }

    heap.Init(&h)

    result := make([][]int, 0, k)
    for i := 0; i < k; i++ {
        result = append(result, heap.Pop(&h).(P).point)
    }

    return result
}
type P struct {
	point []int
	dist  int
}

type minHeap []P

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(P))
}

func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func dist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
