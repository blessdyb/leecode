package main

func sumZero(n int) []int {
	ret := []int{}
	if n%2 == 1 {
		ret = append(ret, 0)
		n = n - 1
	}
	for i := 0; i < n/2; i++ {
		ret = append(ret, i+1)
		ret = append(ret, -i-1)
	}
	return ret
}
