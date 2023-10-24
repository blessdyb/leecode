package main

func equalFrequency(word string) bool {
	hashmap := map[byte]int{}
	for i := 0; i < len(word); i++ {
		hashmap[word[i]]++
	}
	if len(hashmap) == 1 {
		return true
	}
	valuesHashmap := map[int]int{}
	values := []int{}
	for _, v := range hashmap {
		valuesHashmap[v]++
	}
	for v, _ := range valuesHashmap {
		values = append(values, v)
	}
	if len(valuesHashmap) == 1 && values[0] == 1 {
		return true
	} else if len(valuesHashmap) != 2 {
		return false
	}
	if values[0] > values[1] {
		values[0], values[1] = values[1], values[0]
	}
	return valuesHashmap[1] == 1 || (values[1]-values[0] == 1 && valuesHashmap[values[1]] == 1)
}
