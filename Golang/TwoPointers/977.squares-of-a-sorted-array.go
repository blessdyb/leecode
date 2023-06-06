package main

/**
 * Since the array provided is sorted in non-decending order, we can utilize two pointers to
 * tranverse the entire list: one moving forward and the other moving backward. Additionally,
 * we use a third pointer to populate an empty result array from the end to the beginning. This
 * approach ensures that we compare and obtain the largest squared number in the current range.
 * We continue comparing the squared values of the two position pointers until they converge.
 */
func sortedSquares(nums []int) []int {
	length := len(nums)
	ret := make([]int, length, length)
	i, j, k := 0, length-1, length-1
	for i <= j {
		i2 := nums[i] * nums[i]
		j2 := nums[j] * nums[j]
		if i2 <= j2 {
			ret[k] = j2
			j--
		} else {
			ret[k] = i2
			i++
		}
		k--
	}
	return ret
}
