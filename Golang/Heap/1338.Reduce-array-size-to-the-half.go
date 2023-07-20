package main

func minSetSize(arr []int) int {
	heap := [][2]int{}
	hashmap := map[int]int{}
	for _, item := range arr {
		hashmap[item]++
	}
	for key, value := range hashmap {
		heap = append(heap, [2]int{key, value})
		index := len(heap) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && heap[parentIndex][1] < heap[index][1] {
			heap[parentIndex], heap[index] = heap[index], heap[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	}
	ret := 0
	target := len(arr) / 2
	for target > 0 {
		target -= heap[0][1]
		heap[0] = heap[len(heap)-1]
		ret++
		parentIndex := 0
		left := 2*parentIndex + 1
		for left < len(heap)-1 {
			child := left
			if child+1 < len(heap)-1 && heap[child+1][1] > heap[child][1] {
				child++
			}
			if heap[parentIndex][1] < heap[child][1] {
				heap[parentIndex], heap[child] = heap[child], heap[parentIndex]
				parentIndex = child
				left = 2*parentIndex + 1
			} else {
				break
			}
		}
		heap = heap[:len(heap)-1]
	}
	return ret
}
