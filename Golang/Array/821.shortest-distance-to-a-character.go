package main

func shortestToChar(s string, c byte) []int {
	indexes := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			indexes = append(indexes, i)
		}
	}
	ret := make([]int, len(s))
	if len(indexes) == 1 {
		for i := 0; i < len(s); i++ {
			ret[i] = indexes[0] - i
			if ret[i] < 0 {
				ret[i] *= -1
			}
		}
	} else {
		left, right := 0, 1
		for i := 0; i < len(s); i++ {
			if i <= indexes[left] {
				ret[i] = indexes[left] - i
			} else if i > indexes[right] {
				ret[i] = i - indexes[right]
			} else if i > indexes[left] && i < indexes[right] {
				if i-indexes[left] < indexes[right]-i {
					ret[i] = i - indexes[left]
				} else {
					ret[i] = indexes[right] - i
				}
			} else if indexes[right] == i {
				ret[i] = 0
				if right < len(indexes)-1 {
					left, right = right, right+1
				}
			}
		}
	}
	return ret
}
