/*
739. 每日温度
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/
package main

func dailyTemperatures(T []int) []int {

	stack := make([][2]int, 0)
	res := make([]int, len(T))
	stack = append(stack, [2]int{0, 0})
	res[len(T)-1] = 0

	stack = append(stack, [2]int{T[len(T)-1], len(T) - 1})

	for i := len(T) - 2; i >= 0; i-- {
		t := T[i]

		for len(stack) > 0 {
			top := stack[len(stack)-1][0]
			topI := stack[len(stack)-1][1]

			if t == top {
				if t < stack[len(stack)-2][0] {
					res[i] = stack[len(stack)-2][1] - i
					stack[len(stack)-1][1] = i
					break
				} else {
					res[i] = 0
					stack[len(stack)-1][1] = i
					break
				}
			}
			if t < top {

				stack = append(stack, [2]int{t, i})
				res[i] = topI - i
				break
			} else {
				stack = stack[0 : len(stack)-1] // 弹出栈顶元素
				if len(stack) == 1 {
					stack = append(stack, [2]int{t, i}) // 当栈顶元素为0 时插入当前元素
					res[i] = 0
					break
				}
			}
		}
	}

	return res
}
