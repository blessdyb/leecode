package main

func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		total := numbers[l] + numbers[r]
		if total == target {
			return []int{l + 1, r + 1}
		} else if total < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
