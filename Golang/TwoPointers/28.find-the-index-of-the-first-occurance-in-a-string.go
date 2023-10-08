package main

func strStr(haystack string, needle string) int {
	ret := -1
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if len(haystack)-i < len(needle) {
			break
		}
		if haystack[i+j] == needle[j] {
			j++
			if j == len(needle) {
				ret = i
				break
			}
		} else {
			i++
			j = 0
		}
	}
	return ret
}
