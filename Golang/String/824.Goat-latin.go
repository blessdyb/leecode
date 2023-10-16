package main

import "strings"

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	ret := make([]string, len(words))
	for index, word := range words {
		temp := ""
		if word[0] == 'A' || word[0] == 'E' || word[0] == 'I' || word[0] == 'O' || word[0] == 'U' || word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' {
			temp = word + "ma"
		} else {
			temp = word[1:] + word[:1] + "ma"
		}
		for j := 0; j < index+1; j++ {
			temp += "a"
		}
		ret[index] = temp
	}
	return strings.Join(ret, " ")
}
