package main

import (
	"sort"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	sort.Slice(words, func(a, b int) bool {
		return words[a][len(words[a])-1] < words[b][len(words[b])-1]
	})
	ret := []string{}
	for _, word := range words {
		ret = append(ret, word[:len(word)-1])
	}
	return strings.Join(ret, " ")
}
