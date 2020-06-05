/**
560. 和为K的子数组
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/
package main

import "fmt"

/*
方法二：前缀和 + 哈希表优化
思路和算法

我们可以基于方法一利用数据结构进行进一步的优化，我们知道方法一的瓶颈在于对每个 i，我们需要枚举所有的 j 来判断是否符合条件，这一步是否可以优化呢？答案是可以的。

我们定义 pre[i] 为 [0..i] 里所有数的和，则 pre[i] 可以由 pre[i−1] 递推而来，即：

pre[i]=pre[i−1]+nums[i]

那么[j..i] 这个子数组和为 k这个条件我们可以转化为

pre[i]−pre[j−1]==k

简单移项可得符合条件的下标 j 需要满足

pre[j−1]==pre[i]−k

所以我们考虑以 i 结尾的和为 k 的连续子数组个数时只要统计有多少个前缀和为 pre[i]−k 的 pre[j] 即可。我们建立哈希表 mp，以和为键，出现次数为对应的值，记录 pre[i] 出现的次数，从左往右边更mp 边计算答案，那么以 i 结尾的答案 mp[pre[i]−k] 即可在 O(1)时间内得到。最后的答案即为所有下标结尾的和为 k 的子数组个数之和。

需要注意的是，从左往右边更新边计算的时候已经保证了mp[pre[i]−k] 里记录的 pre[j] 的下标范围是0≤j≤i 。同时，由于pre[i] 的计算只与前一项的答案有关，因此我们可以不用建立 pre 数组，直接用 pre 变量来记录 pre[i−1] 的答案即可。

。
*/
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	hashMap := make(map[int]int, 0)
	// fmt.Println(hashMap[5])
	var sum, res int
	hashMap[0] = 1
	for _, v := range nums {
		sum += v
		if _, ok := hashMap[sum-k]; ok {
			res += hashMap[sum-k]
		}
		hashMap[sum]++ //map已经初始化，但还没赋值时有默认值吗？
	}
	return res
}

func subarraySum1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	var sum, res int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if k == sum {
				res++
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	r := subarraySum(nums, k)
	fmt.Println(r)
}
