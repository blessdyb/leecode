package main

func numJewelsInStones(jewels string, stones string) int {
	hashmap := map[rune]bool{}
	for _, j := range jewels {
		hashmap[j] = true
	}
	ret := 0
	for _, s := range stones {
		if hashmap[s] {
			ret++
		}
	}
	return ret
}
