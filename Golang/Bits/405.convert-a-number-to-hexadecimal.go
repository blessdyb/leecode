package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hexChars := "0123456789abcdef"
	result := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		result[i] = hexChars[num&0xf]
		num = num >> 4
	}
	for i := 0; i < 8; i++ {
		if result[i] != '0' {
			result = result[i:]
			break
		}
	}
	return string(result)
}
