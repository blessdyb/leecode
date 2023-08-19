package main

func createTargetArray(nums []int, index []int) []int {
	ret := []int{}
	for i := 0; i < len(index); i++ {
		ret = append(append(append([]int{}, ret[:index[i]]...), nums[i]), ret[index[i]:]...)
	}
	return ret
}
