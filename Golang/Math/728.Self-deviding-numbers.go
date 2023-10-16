package main

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for i := left; i <= right; i++ {
		j := i
		flag := true
		for j > 0 {
			digit := j % 10
			if digit == 0 {
				flag = false
				break
			}
			if i%digit != 0 {
				flag = false
				break
			}
			j /= 10
		}
		if flag {
			ret = append(ret, i)
		}
	}
	return ret
}
