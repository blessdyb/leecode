package main

func uniqueMorseRepresentations(words []string) int {
	ret := map[string]int{}
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, word := range words {
		temp := ""
		for i := 0; i < len(word); i++ {
			temp += codes[word[i]-'a']
		}
		ret[temp]++
	}

	return len(ret)
}
