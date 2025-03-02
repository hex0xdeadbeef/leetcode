package middle

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/?envType=study-plan-v2&envId=leetcode-75

func longestSubarray(nums []int) int {

	var (
		k    = 1
		l, r int

		placedOnesCount int
		q               = newQueue(k)

		maxOnesSequenceLength int
	)

	for r = 0; r < len(nums); r++ {

		switch nums[r] {
		case 0:
			switch placedOnesCount {
			case 0:
				q.push(r)
				placedOnesCount++
			case k:
				l = q.pop() + 1
				q.push(r)
			default:
				q.push(r)
				placedOnesCount++
			}
		case 1:
		}

		if rng := r - l + 1; rng > maxOnesSequenceLength {
			maxOnesSequenceLength = rng
		}
	}

	return maxOnesSequenceLength - 1

}

type queue struct {
	body []int
}

func newQueue(l int) *queue {
	q := queue{body: make([]int, 0, l)}

	return &q
}

func (q *queue) push(n int) {
	q.body = append(q.body, n)
}

func (q *queue) pop() int {
	el := q.body[0]
	q.body = q.body[1:]

	return el
}

func (q *queue) peek() int {
	return q.body[0]
}
