package main

func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		spaces := 0
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == ' ' {
				spaces++
			}
		}
		if spaces+1 > max {
			max = spaces + 1
		}
	}
	return max
}
