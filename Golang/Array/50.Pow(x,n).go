package main

func myPow(x float64, n int) float64 {
	cache := map[int]float64{}
	var helper func(n int) float64
	helper = func(n int) float64 {
		if value, ok := cache[n]; ok {
			return value
		}
		if n == 0 {
			return 1
		} else if n < 0 {
			return 1 / helper(n*-1)
		} else if n == 1 {
			return x
		}
		ret := helper(n/2) * helper(n/2)
		if n%2 == 1 {
			ret *= x
		}
		cache[n] = ret
		return ret
	}
	return helper(n)
}
