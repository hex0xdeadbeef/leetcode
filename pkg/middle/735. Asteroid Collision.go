package middle

type stack struct {
	body []int
}

func newStack(size int) *stack {
	s := stack{body: make([]int, 0, size)}
	return &s
}

func (s *stack) Push(n int) {
	s.body = append(s.body, n)
}

func (s *stack) Peek() int {
	return s.body[len(s.body)-1]
}

func (s *stack) Pop() int {
	elem := s.body[0]
	s.body = s.body[:len(s.body)-1]

	return elem
}

func (s *stack) Elements() []int {
	return s.body
}

func (s *stack) IsNotEmpty() bool {
	return len(s.body) != 0
}

func asteroidCollision(asteroids []int) []int {
	var (
		stack = newStack(len(asteroids))
	)

	for _, a := range asteroids {

		if a > 0 {
			stack.Push(a)
			continue
		}

		if a < 0 {
			collision(stack, a)
		}

	}

	return stack.Elements()
}

func collision(s *stack, negative int) {
	for s.IsNotEmpty() {

		if s.Peek() < 0 {
			break
		}

		if abs(negative) == s.Peek() {
			s.Pop()
			return
		}

		if abs(negative) <= s.Peek() {
			return
		}

		s.Pop()
	}

	s.Push(negative)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}

	return n
}
