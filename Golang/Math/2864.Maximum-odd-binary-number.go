package main

func maximumOddBinaryNumber(s string) string {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	hashmap['1']--
	result := make([]byte, len(s))
	result[len(result)-1] = '1'
	index := 0
	for hashmap['1'] > 0 {
		result[index] = '1'
		hashmap['1']--
		index++
	}
	for hashmap['0'] > 0 {
		result[index] = '0'
		hashmap['0']--
		index++
	}
	return string(result)
}
