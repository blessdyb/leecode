package main

func makeGood(s string) string {
	diff := int('a' - 'A')
	chars := []byte(s)
	for i := 0; i < len(chars) && len(chars) > 1; {
		if i+1 < len(chars) && (int(chars[i])-int(chars[i+1]) == diff || int(chars[i])-int(chars[i+1]) == -diff) {
			chars = append(chars[:i], chars[i+2:]...)
			i = 0
			continue
		}
		i++
	}
	return string(chars)
}
