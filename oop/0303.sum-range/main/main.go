package main

/**
303. 区域和检索 - 数组不可变
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。
*/

/*type NumArray struct {

}

func Constructor(nums []int) NumArray {

}

func (this *NumArray) SumRange(i int, j int) int {

}
*/

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

/**
提前计算出前 i 个数的和，即 a[i] 等于区间 [1,i] 的和，消耗时间为 O(n)

区间 [i,j] 的和求法为：sum(i,j) = a[j] - a[i-1]，每次查询消耗时间为 O(1)，k 次查询时间复杂度为 O(k)

总的来说，时间复杂度为 O(n+k)

作者：elliotxx
链接：https://leetcode-cn.com/problems/range-sum-query-immutable/solution/48mssan-chong-fang-fa-de-go-shi-xian-by-elliotxx/
*/
type NumArray struct {
	presum []int // 存储 [0,i] 的和
}

func Constructor(nums []int) NumArray {
	a := NumArray{}
	a.presum = append(a.presum, 0) // 初始化 presum 数组
	for i := 1; i <= len(nums); i++ {
		t := a.presum[i-1] + nums[i-1]
		a.presum = append(a.presum, t)
	}
	return a
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.presum[j+1] - this.presum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
