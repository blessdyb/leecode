package main

func greatestLetter(s string) string {
	hashmap := map[byte]int{}
	var ret string
	for i := 0; i < len(s); i++ {
		temp := s[i]
		if 'a' <= temp && temp <= 'z' {
			temp = temp - ('a' - 'A')
			hashmap[temp] |= 1
		} else {
			hashmap[temp] |= 2
		}
		if hashmap[temp]^3 == 0 {
			if ret == "" || ret < string(temp) {
				ret = string(temp)
			}
		}
	}
	return ret
}
