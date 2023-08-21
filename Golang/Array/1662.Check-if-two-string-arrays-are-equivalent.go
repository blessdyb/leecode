package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	x, y := "", ""
	for _, w := range word1 {
		x += w
	}
	for _, w := range word2 {
		y += w
	}
	return x == y
}
