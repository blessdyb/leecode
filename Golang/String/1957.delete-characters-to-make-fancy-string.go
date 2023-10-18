package main

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	bytes := []byte{s[0], s[1]}
	for i := 2; i < len(s); i++ {
		if s[i] != bytes[len(bytes)-1] || s[i] != bytes[len(bytes)-2] {
			bytes = append(bytes, s[i])
		}
	}
	return string(bytes)
}
