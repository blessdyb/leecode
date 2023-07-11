package main

func topKFrequent(nums []int, k int) []int {
	ret := []int{}
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	heap := [][2]int{}
	for k, v := range frequency {
		heap = append(heap, [2]int{k, v})
		index := len(heap) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && heap[parentIndex][1] < heap[index][1] {
			heap[index], heap[parentIndex] = heap[parentIndex], heap[index]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	}
	for k > 0 {
		ret = append(ret, heap[0][0])
		heap[0] = heap[len(heap)-1]
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < len(heap)-1 {
			child := left
			if child+1 < len(heap)-1 && heap[child][1] < heap[child+1][1] {
				child = left + 1
			}
			if heap[parentIndex][1] < heap[child][1] {
				heap[child], heap[parentIndex] = heap[parentIndex], heap[child]
				parentIndex = child
				left = parentIndex*2 + 1
			} else {
				break
			}
		}
		heap = heap[:len(heap)-1]
		k--
	}
	return ret
}
