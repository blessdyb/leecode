package main

import "sort"

func checkStrings(s1 string, s2 string) bool {
	getSubSequence := func(s string, index int) string {
		bytes := []byte{}
		for i := index; i < len(s); i += 2 {
			bytes = append(bytes, s[i])
		}
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		return string(bytes)
	}
	a, b, c, d := getSubSequence(s1, 0), getSubSequence(s1, 1), getSubSequence(s2, 0), getSubSequence(s2, 1)
	return a == c && b == d
}
