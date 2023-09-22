package main

func rearrangeArray(nums []int) []int {
	ret, pos, neg := []int{}, []int{}, []int{}
	for _, num := range nums {
		if num > 0 {
			pos = append(pos, num)
		} else {
			neg = append(neg, num)
		}
	}
	for i := 0; i < len(pos); i++ {
		ret = append(ret, pos[i], neg[i])
	}
	return ret
}
