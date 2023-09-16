package main

func replaceDigits(s string) string {
	chars := []byte{}
	shift := func(c byte, i int) byte {
		return c + byte(i)
	}
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			chars[i+1] = shift(chars[i+1], int(chars[i+1]-'0'))
			i++
		}
		i++
	}
	return string(chars)
}
