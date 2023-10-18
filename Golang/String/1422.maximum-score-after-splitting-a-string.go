package main

func maxScore(s string) int {
	ret := 0
	scores := func(index int) (int, int) {
		left, right := 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == '0' && i < index {
				left++
			}
			if s[i] == '1' && i >= index {
				right++
			}
		}
		return left, right
	}
	for i := 1; i <= len(s)-1; i++ {
		s1, s2 := scores(i)
		if ret < s1+s2 {
			ret = s1 + s2
		}
	}
	return ret
}
