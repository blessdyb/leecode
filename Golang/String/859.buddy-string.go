package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		hashmap := map[byte]int{}
		for i := 0; i < len(s); i++ {
			hashmap[s[i]]++
			if hashmap[s[i]] > 1 {
				return true
			}
		}
		return false
	}
	diff := 0
	a, b := []byte{}, []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff++
			a = append(a, s[i])
			b = append(b, goal[i])
		}
		if diff > 2 {
			return false
		}
	}
	if diff == 1 {
		return false
	}
	a[0], a[1] = a[1], a[0]
	return string(a) == string(b)
}
