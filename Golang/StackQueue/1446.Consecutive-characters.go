package main

func maxPower(s string) int {
	bytes := []byte{s[0]}
	ret := 1
	for i := 1; i < len(s); i++ {
		if s[i] == bytes[len(bytes)-1] {
			bytes = append(bytes, s[i])
		} else {
			if len(bytes) > ret {
				ret = len(bytes)
			}
			bytes = []byte{s[i]}
		}
	}
	if len(bytes) > ret {
		ret = len(bytes)
	}
	return ret
}
