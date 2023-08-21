package main

func truncateSentence(s string, k int) string {
	ret := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			k--
		}
		if k == 0 {
			break
		}
		ret = append(ret, s[i])
	}
	return string(ret)
}
