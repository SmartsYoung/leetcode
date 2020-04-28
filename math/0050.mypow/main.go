package _050_mypow

import "math"

func myPow(x float64, n int) float64 {
	res := math.Pow(x, float64(n))
	return res
}
