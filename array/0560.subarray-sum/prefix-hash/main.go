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
