package iterative

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
链接：https://leetcode-cn.com/problems/generate-parentheses

*/
//闭合数
func generateParenthesis(n int) []string {
	return closed(n)
}

func closed(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			leftCombination := closed(i)
			rightCombination := closed(n - 1 - i)
			for _, left := range leftCombination {
				for _, right := range rightCombination {
					res = append(res, "("+left+")"+right)
				}
			}
			// （）
			//  0 = ""
			//  1 = "" ()
			//  2 = "" ()()  (())
			//  3 = "" (()())  ()(())  (())() ((()))
		}
	}
	return res
}
