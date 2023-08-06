package main

func sortArray(nums []int) []int {
	heap := []int{}
	for _, num := range nums {
		heap = append(heap, num)
		index := len(heap) - 1
		parent := (index - 1) / 2
		for index >= 0 && heap[parent] > heap[index] {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
			parent = (index - 1) / 2
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = heap[0]
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		index := 0
		left := 2*index + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && heap[child+1] < heap[child] {
				child++
			}
			if heap[child] < heap[index] {
				heap[child], heap[index] = heap[index], heap[child]
				index = child
				left = 2*index + 1
			} else {
				break
			}
		}
	}
	return nums
}
