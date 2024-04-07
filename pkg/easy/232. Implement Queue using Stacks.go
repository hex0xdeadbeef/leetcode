package easy

// https://leetcode.com/problems/implement-queue-using-stacks/

const (
	initialSize = 512
)

type myQueue struct {
	l, r int
	body []int
}

func constructor() myQueue {
	q := myQueue{body: make([]int, initialSize), l: -1, r: 0}
	return q
}

func (this *myQueue) Push(x int) {
	if this.Empty() {
		this.l = 0
	}

	this.body[this.r] = x
	this.r++

}

func (this *myQueue) Pop() int {
	res := this.body[this.l]
	this.l++

	if this.l == this.r {
		this.l = -1
		this.r = 0
	}

	return res
}

func (this *myQueue) Peek() int {
	return this.body[this.l]
}

func (this *myQueue) Empty() bool {
	return this.l < 0
}
