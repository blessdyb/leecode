package main

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(a, b int) bool {
		return len(dictionary[a]) < len(dictionary[b])
	})
	words := strings.Split(sentence, " ")
	ret := []string{}
	for _, word := range words {
		afterReplace := word
		for _, dict := range dictionary {
			if len(word) >= len(dict) && word[:len(dict)] == dict {
				afterReplace = dict
				break
			}
		}
		ret = append(ret, afterReplace)
	}
	return strings.Join(ret, " ")
}
