package main

func checkDistances(s string, distance []int) bool {
	hashmap := map[byte]int{}
	hashmap[s[0]] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[0] && distance[int(s[i]-'a')] != i-1 {
			return false
		} else if hashmap[s[i]] == 0 {
			hashmap[s[i]] = i
		} else if distance[int(s[i]-'a')] != i-hashmap[s[i]]-1 {
			return false
		}
	}
	return true
}
