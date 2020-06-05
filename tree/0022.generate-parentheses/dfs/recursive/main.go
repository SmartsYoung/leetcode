package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n < 1 {
		return res
	}
	dfs("", n, n, &res)
	return res
}

func dfs(str string, left int, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	}

	if left > right { //剪枝
		return
	}
	if left > 0 {
		dfs(str+"(", left-1, right, res) //左树遍历
	}

	if right > 0 {
		dfs(str+")", left, right-1, res) //右树遍历
	}
}

// dfs

/*回溯算法，回溯跳出条件就是左右括号都已经排完的情况。
括号成对存在，先有左括号再有右括号，所以只有右括号的数量小于左括号才进行右括号的添加。
最后如果右括号的数量等于0，表示右括号已经排完了，同时意味着左括号也排完了。*/

func generateParenthesis1(n int) []string {
	// 使用new方法开辟内存返回内存地址
	res := new([]string)

	backtracking(n, n, "", res)

	return *res
}

func backtracking(left, right int, tmp string, res *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		*res = append(*res, tmp)
		return
	}

	// 生成左括号
	if left > 0 {
		backtracking(left-1, right, tmp+"(", res)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backtracking(left, right-1, tmp+")", res)
	}
}
