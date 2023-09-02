package main

func toLowerCase(s string) string {
	letters := []byte(s)
	for i := 0; i < len(letters); i++ {
		if letters[i] >= 'A' && letters[i] <= 'Z' {
			letters[i] += 'a' - 'A'
		}
	}
	return string(letters)
}
