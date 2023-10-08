package main

import "strings"

func wordPattern(pattern string, str string) bool {
	h1 := map[byte]string{}
	h2 := map[string]byte{}
	words := strings.Split(str, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if _, ok := h1[pattern[i]]; !ok {
			if _, ok2 := h2[words[i]]; ok2 {
				return false
			}
			h1[pattern[i]] = words[i]
			h2[words[i]] = pattern[i]
		}
		if h1[pattern[i]] != words[i] {
			return false
		}
	}
	return true
}
