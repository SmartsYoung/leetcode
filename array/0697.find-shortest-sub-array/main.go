/**
697. 数组的度
给定一个非空且只包含非负数的整数数组 nums, 数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

示例 1:

输入: [1, 2, 2, 3, 1]
输出: 2
解释:
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.
示例 2:

输入: [1,2,2,3,1,4,2]
输出: 6
注意:

nums.length 在1到50,000区间范围内。
nums[i] 是一个在0到49,999范围内的整数。
*/

package main

import "fmt"

var start int
var end int

func minx(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func findShortestSubArray(nums []int) int {
	var max int //若定义为全局变量，则在某些情况下会出错
	var vcount int

	hashMap := make(map[int]int, 0)
	maxMap := make(map[int]int, 0)

	var min int = len(nums)

	for _, v := range nums {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = 0
		}
		hashMap[v] += 1

		if max < hashMap[v] {
			max = hashMap[v]
		}
	}

	for i, v := range nums {
		if nums[i] == v && max == hashMap[v] {
			maxMap[v] = max
		}
	}

	for val, _ := range maxMap {
		for i := 0; i < len(nums); i++ {
			if nums[i] == val {
				if vcount == 0 {
					start = i
				}
				vcount++
				if vcount == max {
					end = i
					vcount = 0
					break
				}
			}
		}
		mi := end - start + 1
		min = minx(mi, min)
	}
	return min
}

func main() {
	arr := []int{1, 2, 2, 3, 1}
	arr1 := [7]int{1, 2, 2, 3, 1, 4, 2}
	res := findShortestSubArray(arr)
	fmt.Println(res)
	fmt.Println(findShortestSubArray(arr1[:]))
	fmt.Println(findShortestSubArray([]int{3, 3, 4}))
}

//
func findShortestSubArray1(nums []int) int {
	// 边界判断，如果是0就返回0就OK
	if len(nums) == 0 {
		return 0
	}
	// 创建一个辅助哈希，哈希v是一个切片
	r := make(map[int][]int)
	for k, v := range nums {
		if _, ok := r[v]; ok {
			r[v] = append(r[v], k) //  切片的类型为数组的索引值
		} else {
			// 因为不知道len，所以这个slice最开始创建要使用len == 1
			a := make([]int, 1, 1)
			a[0] = k
			r[v] = a
		}
	}

	// 1 是数量 2 是index的距离 result 就是最后的一个输出指针，
	// 所有的值跟result不断的比较，然后最终result就是真实要输出的值。
	// 因为只需要两个值，并且不存在数组的复制，所以直接使用数组就OK了，
	// 没必要使用slice。
	result := [2]int{0, 0}

	for _, v := range r {

		// 先比len，len大的取胜，如果len一样，那么下标差小的取胜。
		if result[0] < len(v) || result[0] == len(v) && result[1] > v[len(v)-1]-v[0] {
			result = [2]int{len(v), v[len(v)-1] - v[0]}
		}

	}

	return result[1] + 1
}

/*作者：googege
链接：https://leetcode-cn.com/problems/degree-of-an-array/solution/goyu-yan-ban-ben-shi-yong-yi-ge-mapzuo-fu-zhu-by-g/*/
