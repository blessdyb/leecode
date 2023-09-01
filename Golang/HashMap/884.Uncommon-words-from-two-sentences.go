package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	arr1 := strings.Split(s1, " ")
	arr2 := strings.Split(s2, " ")
	h1 := map[string]int{}
	h2 := map[string]int{}

	for _, item := range arr1 {
		h1[item] |= 1
		h2[item]++
	}
	for _, item := range arr2 {
		h1[item] |= 2
		h2[item]++
	}
	ret := []string{}
	for k, v := range h1 {
		if v != 3 && h2[k] == 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
