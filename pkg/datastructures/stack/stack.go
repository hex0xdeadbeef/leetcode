package stack

type Stack struct {
	body []int
}

func NewStack(size int) *Stack {
	s := Stack{body: make([]int, 0, size)}
	return &s
}

func (s *Stack) Push(n int) {
	s.body = append(s.body, n)
}

func (s *Stack) Peek() int {
	return s.body[len(s.body)-1]
}

func (s *Stack) Pop() int {
	elem := s.body[0]
	s.body = s.body[:len(s.body)-1]

	return elem
}

func (s *Stack) Elements() []int {
	return s.body
}

func (s *Stack) IsNotEmpty() bool {
	return len(s.body) != 0
}
