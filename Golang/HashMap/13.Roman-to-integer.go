package main

func romanToInt(s string) int {
	hashmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ret := 0
	for i := 0; i < len(s); {
		if s[i] == 'I' && i+1 < len(s) {
			if s[i+1] == 'V' {
				ret += 4
				i += 2
				continue
			} else if s[i+1] == 'X' {
				ret += 9
				i += 2
				continue
			}
		} else if s[i] == 'X' && i+1 < len(s) {
			if s[i+1] == 'L' {
				ret += 40
				i += 2
				continue
			} else if s[i+1] == 'C' {
				ret += 90
				i += 2
				continue
			}
		} else if s[i] == 'C' && i+1 < len(s) {
			if s[i+1] == 'D' {
				ret += 400
				i += 2
				continue
			} else if s[i+1] == 'M' {
				ret += 900
				i += 2
				continue
			}
		}
		ret += hashmap[s[i]]
		i++
	}
	return ret
}
