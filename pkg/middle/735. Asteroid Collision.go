package middle

import . "leetcode/pkg/datastructures/stack"

func asteroidCollision(asteroids []int) []int {
	var (
		stack = NewStack(len(asteroids))
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

func collision(s *Stack, negative int) {
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
