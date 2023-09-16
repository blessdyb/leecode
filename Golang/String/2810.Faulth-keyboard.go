package main

func finalString(s string) string {
	ret := []byte{}
	reverse := func(str []byte) {
		for i, j := 0, len(str)-1; i < j; {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			reverse(ret)
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
