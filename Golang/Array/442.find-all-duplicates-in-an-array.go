package main

// nums are in [1, n] and appears [0, 1, 2] times
// https://www.mindluster.com/lesson/163700
func findDuplicates(nums []int) []int {
	ret := []int{}
	for _, num := range nums {
		numberIndexAfterSorted := num
		if num < 0 {
			numberIndexAfterSorted *= -1
		}
		nums[numberIndexAfterSorted-1] *= -1
		if nums[numberIndexAfterSorted-1] > 0 {
			ret = append(ret, numberIndexAfterSorted)
		}
	}
	return ret
}
