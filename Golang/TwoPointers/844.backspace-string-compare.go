package main

func backspaceCompare(s string, t string) bool {
	return removeBackspace(s) == removeBackspace(t)
}

// Use two pointers to in-place sanitize the string
func removeBackspace(str string) string {
	chars := []rune(str)
	length := len(chars)
	i, j := 0, 0
	for j < length {
		if chars[j] == '#' {
			chars[i] = chars[j]
			i++
		} else {
			i--
		}
		if i < 0 {
			i = 0
		}
		j++
	}
	return string(chars[:i])
}
