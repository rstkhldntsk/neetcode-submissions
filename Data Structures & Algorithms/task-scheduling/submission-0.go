func leastInterval(tasks []byte, n int) int {
	count := [26]int{}
	for _, task := range tasks {
		count[task-'A']++
	}
	
	h := make([]int, 0)
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			h = insert(h, count[i])
		}
	}
	
	time := 0
	q := make([][2]int, 0)
	for len(h) > 0 || len(q) > 0 {
		time++
		var cnt int
		if len(h) == 0 {
			time = q[0][1]
		} else {
			h, cnt = pop(h)
			cnt--
			if cnt > 0 {
				q = append(q, [2]int{cnt, time + n})
			}
		}
		if len(q) > 0 && q[0][1] == time {
			h = insert(h, q[0][0])
			q = q[1:]
		}
	}
	return time
}

func insert(h []int, i int) []int {
	h = append(h, i)
	siftUp(h, len(h)-1)
	return h
}

func pop(h []int) ([]int, int) {
	e := h[0]
	h[0], h[len(h)-1] = h[len(h)-1], h[0]
	h = h[:len(h)-1]
	siftDown(h, 0)
	return h, e
}

func siftUp(h []int, i int) {
	for i >= 0 && h[parent(i)] < h[i] {
		h[i], h[parent(i)] = h[parent(i)], h[i]
		i = parent(i)
	}
}

func siftDown(h []int, i int) {
	for left(i) < len(h) {
		j := i
		if h[left(i)] > h[i] {
			j = left(i)
		}
		if right(i) < len(h) && h[right(i)] > h[j] {
			j = right(i)
		}
		if j == i {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}
