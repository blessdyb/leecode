package main

func restoreString(s string, indices []int) string {
	ret := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ret[indices[i]] = s[i]
	}
	return string(ret)
}
