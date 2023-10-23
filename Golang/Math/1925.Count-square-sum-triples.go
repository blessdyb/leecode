package main

func countTriples(n int) int {
	nums := []int{}
	hashmap := map[int]bool{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i*i)
		hashmap[i*i] = true
	}
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if hashmap[nums[i]+nums[j]] {
				count += 2
			}
		}
	}
	return count
}
