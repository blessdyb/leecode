package main

func lengthOfLongestSubstring(s string) int {
	hashmap := map[byte]bool{}
	max, left, right := 0, 0, 0
	for right < len(s) {
		if !hashmap[s[right]] {
			hashmap[s[right]] = true
			right++
			if max < right-left {
				max = right - left
			}
		} else {
			hashmap[s[left]] = false
			left++
		}
	}
	return max
}
