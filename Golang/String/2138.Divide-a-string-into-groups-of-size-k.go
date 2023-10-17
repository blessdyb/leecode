package main

func divideString(s string, k int, fill byte) []string {
	ret := []string{}
	temp := []byte{}
	for i := 0; i < len(s); i++ {
		temp = append(temp, s[i])
		if len(temp) == k {
			ret = append(ret, string(temp))
			temp = []byte{}
		}
	}
	if len(temp) > 0 {
		for len(temp) < k {
			temp = append(temp, fill)
		}
		ret = append(ret, string(temp))
	}
	return ret
}
