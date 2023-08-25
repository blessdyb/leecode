package main

func prefixCount(words []string, pref string) int {
	ret := 0
	for _, word := range words {
		if len(pref) <= len(word) {
			flag := true
			for i := 0; i < len(pref); i++ {
				if pref[i] != word[i] {
					flag = false
					break
				}
			}
			if flag {
				ret++
			}
		}
	}
	return ret
}
