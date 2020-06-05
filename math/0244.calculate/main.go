/*
224. 基本计算器
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
*/
package main

func calculate(s string) int {
	res := 0
	stack := make([]int, 0, len(s))
	sign := 1
	num := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// 提取 s 中的数字
			num = 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			// 根据前面的记录，进行运算
			res += sign * num
			// 此时 s[i] 已经不是数字了
			// for 语句中，会再＋1，所以这里先 -1
			i--
		case '+':
			// 记录运算符
			sign = 1
		case '-':
			// 记录运算符
			sign = -1
		case '(':
			// 遇到 '(' 就把当前的 res 和 sign 入栈，保存好当前的运行环境
			stack = append(stack, res, sign)
			// 对 res 和 sign 赋予新的值
			res = 0
			sign = 1
		case ')':
			// 遇到 ')' 出栈
			// sign 是与这个 ')' 匹配的 '(' 前的运算符号
			sign = stack[len(stack)-1]
			// temp 是 sign 前的运算结果
			temp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// '(' 与 ')' 之间的运算结果
			//          ↓
			res = sign*res + temp
		}
	}

	return res
}
