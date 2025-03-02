package middle

// https://leetcode.com/problems/min-stack/description/

const (
	defaultSize = 4096
)

type MinStack struct {
	body []int
	mins []int
}

func Constructor() MinStack {

	return MinStack{body: make([]int, 0, defaultSize), mins: make([]int, 0, defaultSize)}
}

func (this *MinStack) Push(val int) {
	this.body = append(this.body, val)

	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
		return
	}

	if val < this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
		return
	}

	this.mins = append(this.mins, this.mins[len(this.mins)-1])
}

func (this *MinStack) Pop() {
	this.body, this.mins = this.body[:len(this.body)-1], this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.body[len(this.body)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
