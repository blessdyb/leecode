package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	getHashmap := func(str string) map[byte]int {
		hashmap := map[byte]int{}
		for i := 0; i < len(str); i++ {
			hashmap[str[i]]++
		}
		return hashmap
	}
	h1 := getHashmap(s)
	h2 := getHashmap(target)
	min := 100
	for k, v := range h2 {
		fmt.Println(h1[k], v)
		if min > (h1[k] / v) {
			min = h1[k] / v
		}
	}
	return min
}
