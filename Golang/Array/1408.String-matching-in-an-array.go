package main

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ret := []string{}
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				ret = append(ret, words[i])
				break
			}
		}
	}
	return ret
}
