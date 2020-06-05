package main

/**
155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

提示：
pop、top 和 getMin 操作总是在 非空栈 上调用。
*/

type MinStack struct {
	s []Node
}
type Node struct {
	d int //data
	m int //current min
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	d := Node{d: x, m: x}
	if len(this.s) > 0 && this.s[len(this.s)-1].m < x {
		d.m = this.s[len(this.s)-1].m
	}
	this.s = append(this.s, d)

}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].d
}

func (this *MinStack) GetMin() int {
	return this.s[len(this.s)-1].m
}
