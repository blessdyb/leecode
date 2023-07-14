package main

func commonChars(words []string) []string {
	hashset := [26]int{}
	for _, c := range words[0] {
		hashset[c-rune('a')]++
	}
	for i := 1; i < len(words); i++ {
		temp := [26]int{}
		for _, c := range words[i] {
			temp[c-rune('a')]++
		}
		for i := 0; i < 26; i++ {
			if hashset[i] > temp[i] {
				hashset[i] = temp[i]
			}
		}
	}
	ret := []string{}
	for idx, value := range hashset {
		for i := 1; i <= value; i++ {
			ret = append(ret, string(rune(idx+int('a'))))
		}
	}
	return ret
}
