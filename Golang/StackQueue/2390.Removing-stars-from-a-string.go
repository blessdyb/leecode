package main

func removeStars(s string) string {
	ret := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			ret = ret[:len(ret)-1]
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
