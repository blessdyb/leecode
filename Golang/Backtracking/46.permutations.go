package main

func permute(nums []int) [][]int {
	ret := [][]int{}
	permutation := []int{}
	var backtracking func([]int)
	backtracking = func(nums []int) {
		if len(nums) == 0 {
			// Note: here we clone permutation instead of using it directly since it's a reference
			ret = append(ret, append([]int{}, permutation...))
		} else {
			for i, num := range nums {
				permutation = append(permutation, num)
				backtracking(append(append([]int{}, nums[:i]...), nums[i+1:]...))
				permutation = permutation[:len(permutation)-1]
			}
		}
	}
	backtracking(nums)
	return ret
}
