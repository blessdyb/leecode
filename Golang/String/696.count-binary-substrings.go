package main

func countBinarySubstrings(s string) int {
	groups := []int{1}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			groups[len(groups)-1] += 1
		} else {
			groups = append(groups, 1)
		}
	}
	ret := 0
	for i := 1; i < len(groups); i++ {
		if groups[i] < groups[i-1] {
			ret += groups[i]
		} else {
			ret += groups[i-1]
		}
	}
	return ret
}

func countBinarySubstrings2(s string) int {
	ret, prev, curr := 0, 0, 1
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			ret += min(prev, curr)
			prev, curr = curr, 1
		}
	}
	return ret + min(prev, curr)
}
