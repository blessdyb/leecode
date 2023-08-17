package main

func similarPairs(words []string) int {
	bits := []int{}
	for _, word := range words {
		temp := 0
		for i := 0; i < len(word); i++ {
			temp = temp | (2 << int(word[i]-'a'))
		}
		bits = append(bits, temp)
	}
	ret := 0
	for i := 0; i < len(bits)-1; i++ {
		for j := i + 1; j < len(bits); j++ {
			if bits[i]^bits[j] == 0 {
				ret++
			}
		}
	}
	return ret
}
