package main

func numberOfLines(widths []int, s string) []int {
	ret := []int{1, 0}
	for i := 0; i < len(s); i++ {
		temp := widths[s[i]-'a']
		if ret[1]+temp > 100 {
			ret[0]++
			ret[1] = temp
		} else {
			ret[1] += temp
		}
	}
	return ret
}
