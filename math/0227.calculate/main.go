package main

import "strconv"

/**
227. 基本计算器 II
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:

输入: "3+2*2"
输出: 7
示例 2:

输入: " 3/2 "
输出: 1
示例 3:

输入: " 3+5 / 2 "
输出: 5
*/

func calculate(s string) int {

	stack := make([]int, 0, len(s))
	op := make([]int, 0, len(s))
	num := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(op) > 0 && op[len(op)-1] > 2 {
				if op[len(op)-1] == 3 {
					stack[len(stack)-1] *= num
				} else {
					stack[len(stack)-1] /= num
				}
				op = op[:len(op)-1]
			} else {
				stack = append(stack, num)
			}
			// 退一格i 上面自动i++
			i--
		case '+':
			op = append(op, 1)
		case '-':
			op = append(op, -1)
		case '*':
			op = append(op, 3)
		case '/':
			op = append(op, 4)
		default:
			// fmt.Println(s[i])
		}
	}
	// fmt.Println(stack, op)
	for len(op) > 0 {
		stack[1] = stack[0] + op[0]*stack[1]
		op = op[1:]
		stack = stack[1:]
	}

	return stack[0]
}

func calculate1(s string) int {
	numstr := "0"
	stack := []int{0}
	op := []byte{'+'}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numstr += string(s[i])
		}

		// 遇到操作符时开始计算上一个操作符的结果， 如  6 - 10 * 2 + 3, 遇到加号时开始计算 10 *2
		//
		if i == len(s)-1 || '+' == s[i] || '-' == s[i] || '*' == s[i] || '/' == s[i] {
			num, err := strconv.Atoi(numstr) // num 是 运算符之前的那个数
			if err != nil {
				break
			}
			devidor := stack[len(stack)-1] // 栈顶元素
			stack = stack[:len(stack)-1]   // 删除栈顶元素
			// 乘除法可以立即运算   加减法不能立即运算，如 5 + 8 * 2 / 4 - 100 / 10 + 3
			// 这个乘号是栈顶的元素， 并不是当前的运算结果
			if '*' == op[len(op)-1] {
				stack = append(stack, devidor*num)
			} else if '/' == op[len(op)-1] {
				stack = append(stack, devidor/num)
			} else if '+' == op[len(op)-1] {
				stack = append(stack, devidor) // 栈顶元素被删除， 重新插入栈顶元素
				stack = append(stack, num)     // 插入当前运算符前面的数
			} else if '-' == op[len(op)-1] {
				stack = append(stack, devidor)
				stack = append(stack, -num)
			}
			op = op[:len(op)-1]   // 栈顶运算符已经被运算，删除
			op = append(op, s[i]) // 插入当前的运算符
			numstr = ""
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
