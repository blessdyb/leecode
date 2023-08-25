package main

func diStringMatch(s string) []int {
	ret := []int{}
	i := 0
	d := len(s)
	for j := 0; j < len(s); j++ {
		if s[j] == 'I' {
			ret = append(ret, i)
			i++
		} else {
			ret = append(ret, d)
			d--
		}
	}
	ret = append(ret, i)
	return ret
}
