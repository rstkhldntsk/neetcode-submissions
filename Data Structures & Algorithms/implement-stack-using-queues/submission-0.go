type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{make([]int,0), make([]int,0)}
}

func (s *MyStack) Push(x int) {
    s.q2 = append(s.q2, x)
    for len(s.q1) > 0 {
        s.q2 = append(s.q2, s.q1[0])
        s.q1 = s.q1[1:]
    }
    s.q1, s.q2 = s.q2, s.q1
}

func (s *MyStack) Pop() int {
	top := s.q1[0]
	s.q1 = s.q1[1:]
	return top
}

func (s *MyStack) Top() int {
	return s.q1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q1) == 0
}
