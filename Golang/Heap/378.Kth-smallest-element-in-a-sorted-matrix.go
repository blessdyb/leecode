package main

func kthSmallest(matrix [][]int, k int) int {
	heap := []int{}
	for _, row := range matrix {
		for _, item := range row {
			if len(heap) < k {
				heap = append(heap, item)
				index := len(heap) - 1
				parent := (index - 1) / 2
				for index >= 0 && heap[parent] < heap[index] {
					heap[parent], heap[index] = heap[index], heap[parent]
					index = parent
					parent = (index - 1) / 2
				}
			} else {
				if heap[0] <= item {
					break
				}
				heap[0] = item
				index := 0
				left := 2*index + 1
				for left < k {
					child := left
					if child+1 < k && heap[child+1] > heap[child] {
						child++
					}
					if heap[child] > heap[index] {
						heap[child], heap[index] = heap[index], heap[child]
						index = child
						left = index*2 + 1
					} else {
						break
					}
				}
			}
		}
	}
	return heap[0]
}
