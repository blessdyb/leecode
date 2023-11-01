package main

// Can't use built-in number types since it may overflow.
func addToArrayForm(num []int, k int) []int {
	ks := []int{}
	for k > 0 {
		ks = append([]int{k % 10}, ks...)
		k /= 10
	}
	carrier := 0
	for i, j := len(num)-1, len(ks)-1; i >= 0 || j >= 0; {
		if i >= 0 && j >= 0 {
			num[i] += carrier + ks[j]
			if num[i] > 9 {
				num[i] -= 10
				carrier = 1
			} else {
				carrier = 0
			}
			i--
			j--
		} else if i >= 0 {
			num[i] += carrier
			if num[i] > 9 {
				num[i] -= 10
				carrier = 1
			} else {
				carrier = 0
			}
			i--
		} else if j >= 0 {
			temp := ks[j] + carrier
			if temp > 9 {
				temp -= 10
				carrier = 1
			} else {
				carrier = 0
			}
			num = append([]int{temp}, num...)
			j--
		}
	}
	if carrier > 0 {
		num = append([]int{carrier}, num...)
	}
	return num
}
