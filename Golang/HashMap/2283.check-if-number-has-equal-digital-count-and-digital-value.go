package main

func digitCount(num string) bool {
	hashmap := map[byte]int{}
	for i := 0; i < len(num); i++ {
		hashmap[num[i]]++
	}
	for i := 0; i < len(num); i++ {
		// Converting between byte and int
		if hashmap[byte(i+'0')] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
