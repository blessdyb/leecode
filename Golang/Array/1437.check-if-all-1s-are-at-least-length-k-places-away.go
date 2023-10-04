package main

func kLengthApart(arr []int, k int) bool {
	lastIndex := -k - 1
	for i := 0; i < len(arr); i++ {
		if nums[i] == 1 {
			if i-lastIndex <= k {
				return false
			} else {
				lastIndex = i
			}
		}
	}
	return true
}
