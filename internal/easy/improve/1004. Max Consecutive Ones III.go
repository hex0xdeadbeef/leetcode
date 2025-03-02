package improve

// https://leetcode.com/problems/max-consecutive-ones-iii/?envType=study-plan-v2&envId=leetcode-75

func LongestOnes(nums []int, k int) int {
	if k == 0 {
		return zeroK(nums)
	}

	return nonZeroK(nums, k)
}

func zeroK(nums []int) int {

	var (
		placedOnesCount       int
		maxOnesSequenceLength int
	)

	for i := 0; i >= 0; i++ {
		if placedOnesCount > maxOnesSequenceLength {
			maxOnesSequenceLength = placedOnesCount
		}

		if i == len(nums) {
			break
		}

		if nums[i] == 0 {
			nums = nums[i+1:]
			placedOnesCount = 0
			i = -1
			continue
		}

		placedOnesCount++
	}

	return maxOnesSequenceLength

}

func nonZeroK(nums []int, k int) int {

	var (
		l, r int

		placedOnesCount int
		q               = NewQueue(k)

		maxOnesSequenceLength int
	)

	for r = 0; r < len(nums); r++ {

		switch nums[r] {
		case 0:
			switch placedOnesCount {
			case 0:
				q.push(r)
				placedOnesCount++
				// fmt.Printf("Tail is %d\n", q.peek())
			case k:
				l = q.pop() + 1
				q.push(r)
				// fmt.Printf("Tail is %d\n", q.peek())
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

	return maxOnesSequenceLength

}

type queue struct {
	body []int
}

func NewQueue(l int) *queue {
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
