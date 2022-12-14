package main

import "math"

// 155-最小栈
type MinStack struct {
	stack  []int
	minStk []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		minStk: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStk[len(this.minStk)-1]
	this.minStk = append(this.minStk, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
