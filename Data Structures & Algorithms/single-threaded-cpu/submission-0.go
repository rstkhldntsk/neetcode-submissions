func getOrder(tasks [][]int) []int {
	result := make([]int, 0)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = append(tasks[i], i)
	}
	sort.Slice(
		tasks, func(i, j int) bool {
			if tasks[i][0] != tasks[j][0] {
				return tasks[i][0] < tasks[j][0]
			}
			return tasks[i][1] < tasks[j][1]
		},
	)
	pq := &PQ{}
	heap.Init(pq)
	time := tasks[0][0]
	i := 0
	for pq.Len() > 0 || i < len(tasks) {
		for i < len(tasks) && time >= tasks[i][0] {
			heap.Push(
				pq,
				Task{enqueueTime: tasks[i][0], processingTime: tasks[i][1], pos: tasks[i][2]},
			)
			i++
		}
		
		if pq.Len() == 0 {
			time = tasks[i][0]
		} else {
			task := heap.Pop(pq).(Task)
			time += task.processingTime
			result = append(result, task.pos)
		}
	}
	
	return result
}

type Task struct {
	enqueueTime    int
	processingTime int
	pos            int
}

type PQ []Task

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	 if pq[i].processingTime == pq[j].processingTime {
      return pq[i].pos < pq[j].pos
  }
  return pq[i].processingTime < pq[j].processingTime
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(Task))
}

func (h *PQ) Pop() any {
	key := *h
	n := len(key)
	x := key[n-1]
	*h = key[0 : n-1]
	return x
}
