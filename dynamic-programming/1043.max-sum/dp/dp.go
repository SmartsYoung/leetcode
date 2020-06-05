package main

import "fmt"

/**
1043. 分隔数组以得到最大和
给出整数数组 A，将该数组分隔为长度最多为 K 的几个（连续）子数组。分隔完成后，每个子数组的中的值都会变为该子数组中的最大值。

返回给定数组完成分隔后的最大和。



示例：

输入：A = [1,15,7,9,2,5,10], K = 3
输出：84
解释：A 变为 [15,15,15,9,10,10,10]


提示：

1 <= K <= A.length <= 500
0 <= A[i] <= 10^6
*/

func maxSumAfterPartitioning(A []int, K int) int {
	dp := make([]int, len(A)+1)

	for i := 0; i <= len(A); i++ {
		j := i - 1

		maxv := dp[i]
		for j >= 0 && i-j <= K {
			maxv = max(maxv, A[j])
			dp[i] = max(dp[i], dp[j]+maxv*(i-j))
			j--
		}
	}
	return dp[len(A)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	arr := []int{1, 15, 7, 9, 2, 5, 10}
	fmt.Println(maxSumAfterPartitioning(arr, 3))

}
