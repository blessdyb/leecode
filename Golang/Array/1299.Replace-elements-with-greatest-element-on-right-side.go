package main

func replaceElements(arr []int) []int {
	n := len(arr)
	hashmap := map[int]int{}
	hashmap[n-1] = -1
	if n >= 2 {
		max := arr[n-1]
		for i := len(arr) - 2; i >= 0; i-- {
			if arr[i+1] > max {
				max = arr[i+1]
			}
			hashmap[i] = max
		}
	}

	for k, v := range hashmap {
		arr[k] = v
	}
	return arr
}
