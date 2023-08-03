package main

func maximumImportance(n int, roads [][]int) int64 {
	hashmap := map[int]int{}
	for _, item := range roads {
		hashmap[item[0]]++
		hashmap[item[1]]++
	}
	heap := [][2]int{}
	for k, v := range hashmap {
		heap = append(heap, [2]int{k, v})
		index := len(heap) - 1
		parent := (index - 1) / 2
		for index >= 0 && heap[parent][1] < heap[index][1] {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
			parent = (index - 1) / 2
		}
	}
	for len(heap) > 0 {
		hashmap[heap[0][0]] = n
		n--
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		parent := 0
		left := 2*parent + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && heap[child+1][1] > heap[child][1] {
				child++
			}
			if heap[parent][1] < heap[child][1] {
				heap[parent], heap[child] = heap[child], heap[parent]
				parent = child
				left = 2*parent + 1
			} else {
				break
			}
		}
	}
	var ret int64 = 0
	for _, item := range roads {
		ret += int64(hashmap[item[0]])
		ret += int64(hashmap[item[1]])
	}
	return ret
}
