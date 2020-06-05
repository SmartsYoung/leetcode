/**
29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/
package main

import "math"

func divide(dividend int, divisor int) int {
	negative := dividend^divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var multiple int32
	var aMultiple = divisor
	for divisor <= dividend {
		d := aMultiple
		m := int32(1)
		for divisor+d<<1 < dividend {
			d <<= 1
			m <<= 1
		}
		multiple += m

		divisor += d
	}
	if multiple < 0 { //处理溢出
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	if negative {
		multiple = -multiple
	}
	if multiple > math.MaxInt32 {
		return math.MaxInt32
	}
	if multiple < math.MinInt32 {
		return math.MinInt32
	}
	return int(multiple)
}

/**
思路：移位倍增
*/
