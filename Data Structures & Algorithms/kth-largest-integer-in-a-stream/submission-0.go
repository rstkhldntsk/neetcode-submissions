type KthLargest struct {
    heap []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		heap = push(heap, nums[i])
		if len(heap) > k {
			heap = pop(heap)
		}
	}
    return KthLargest{
		k: k,
		heap: heap,
	}
}

func (q *KthLargest) Add(val int) int {
	if len(q.heap) < q.k {
		q.heap = push(q.heap, val)
	} else if val > q.heap[0] {
		q.heap = pop(q.heap)
		q.heap = push(q.heap, val)
	}
	
	return q.heap[0]
}

func push(heap []int, val int) []int {
	heap = append(heap, val)
	siftUp(heap, len(heap)-1)
	return heap
}

func pop(heap []int) []int {
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]
	siftDown(heap, 0)
	return heap
}

func siftUp(heap []int, i int) {
	for i > 0 && heap[(i-1)/2] > heap[i] {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
		i = (i-1)/2
	}
}

func siftDown(heap []int, i int) {
	for (i*2)+1 < len(heap) {
		j := i
		if heap[i] > heap[(i*2)+1] {
			j = (i*2)+1
		}
		if (i*2)+2 < len(heap) && heap[(i*2)+2] < heap[j] {
			j = (i*2)+2
		}
		if i == j {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}
