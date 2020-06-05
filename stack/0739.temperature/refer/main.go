package refer

type Temperature struct {
	Value int
	Index int
}

func dailyTemperatures(T []int) []int {
	warmerDay := make([]int, len(T), len(T))
	stack := make([]*Temperature, 0, len(T))

	if len(T) == 0 {
		return nil
	}

	for i, t := range T {
		tStruct := &Temperature{
			Value: t,
			Index: i,
		}
		for len(stack) > 0 && stack[len(stack)-1].Value < t { // 一直循环
			warmerDay[stack[len(stack)-1].Index] = i - stack[len(stack)-1].Index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tStruct)
	}
	return warmerDay
}
