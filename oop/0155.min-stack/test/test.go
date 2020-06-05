package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: nil,
		minStack: nil,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, x)

	if len(this.minStack) > 1 && x > this.minStack[len(this.minStack)-2] {
		this.minStack[len(this.minStack)-1] = this.minStack[len(this.minStack)-2]
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
	this.minStack = this.minStack[0 : len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.GetMin()
	obj.Push(-3)
	obj.GetMin()

	fmt.Println(obj.Top())
	obj.Pop()
}
