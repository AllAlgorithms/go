package fibonacci

// Fibonacci return a fibonacci generator
func Fibonacci() func() int64 {
	var prev int64 = -1
	var res int64 = 1
	return func() int64 {
		res, prev = res+prev, res
		return res
	}
}
