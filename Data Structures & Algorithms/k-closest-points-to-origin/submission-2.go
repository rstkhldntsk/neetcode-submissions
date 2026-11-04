func kClosest(points [][]int, k int) [][]int {
	heapify(points)
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		points, result[i] = pop(points) 
	}
	return result
}

func heapify(heap [][]int) {
	for i := len(heap)/2 - 1; i >= 0; i-- {
		siftDown(heap, i)
	}
}

func siftDown(heap [][]int, i int) {
	for (i*2)+1 < len(heap) {
		j := i
		if dist(heap[(i*2)+1]) < dist(heap[i]) {
			j = (i * 2) + 1
		}
		if (i*2)+2 < len(heap) && dist(heap[(i*2)+2]) < dist(heap[j]) {
			j = (i * 2) + 2
		}
		if j == i {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}

func pop(points [][]int) ([][]int, []int) {
	key := points[0]
	points[0], points[len(points)-1] = points[len(points)-1], points[0]
	points = points[:len(points)-1]
	siftDown(points, 0)
	return points, key
}

func dist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
