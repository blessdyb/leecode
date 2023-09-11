package main

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	ret := []string{}
	for _, word := range words {
		separated := strings.Split(word, string(separator))
		for _, s := range separated {
			if s != "" {
				ret = append(ret, s)
			}
		}
	}
	return ret
}
