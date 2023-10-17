package main

func numDifferentIntegers(word string) int {
	hashmap := map[string]bool{}
	temp := []byte{}
	trimStartingZero := func(digits []byte) string {
		index := -1
		for i := 0; i < len(digits); i++ {
			if digits[i] != '0' {
				index = i
				break
			}
		}
		if index == -1 {
			return "0"
		}
		return string(digits[index:])
	}
	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			temp = append(temp, word[i])
		} else if len(temp) > 0 {
			hashmap[trimStartingZero(temp)] = true
			temp = []byte{}
		}
	}
	if len(temp) > 0 {
		hashmap[trimStartingZero(temp)] = true
	}
	return len(hashmap)
}
