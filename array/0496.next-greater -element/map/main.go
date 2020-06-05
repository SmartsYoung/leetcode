package _map

/*
注意到题目中是没有重复元素的条件，我们可以用一个 map 来存储 nums2 也就是大数组的每个不同值的位置
    存储每个元素的位置
封装一个 findNextBiggerNum() 函数，用来线性查找第一个比指定元素大的元素
遍历 nums1 数组，通过存储好的 map 来划分 nums2 数组，调用 findNextBiggerNum()
代码*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i // 字典顺序
	}

	var ret []int

	for _, n := range nums1 {
		ret = append(ret, findNextBiggerNum(nums2[m[n]:], n))
	}

	return ret
}

func findNextBiggerNum(nums []int, val int) int {
	for _, n := range nums {
		if n > val {
			return n
		}
	}
	return -1
}
