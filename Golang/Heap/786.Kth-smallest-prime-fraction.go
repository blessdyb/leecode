package main

func kthSmallestPrimeFraction(arr []int, k int) []int {
	heap := [][]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if len(heap) < k {
				heap = append(heap, []int{arr[i], arr[j]})
				index := len(heap) - 1
				parent := (index - 1) / 2
				for index >= 0 && (float32(heap[index][0])/float32(heap[index][1]) > float32(heap[parent][0])/float32(heap[parent][1])) {
					heap[index], heap[parent] = heap[parent], heap[index]
					index = parent
					parent = (index - 1) / 2
				}
			} else if float32(heap[0][0])/float32(heap[0][1]) > float32(arr[i])/float32(arr[j]) {
				heap[0] = []int{arr[i], arr[j]}
				index := 0
				left := 2*index + 1
				for left < k {
					child := left
					if child+1 < k && (float32(heap[child+1][0])/float32(heap[child+1][1]) > float32(heap[child][0])/float32(heap[child][1])) {
						child++
					}
					if float32(heap[child][0])/float32(heap[child][1]) > float32(heap[index][0])/float32(heap[index][1]) {
						heap[index], heap[child] = heap[child], heap[index]
						index = child
						left = 2*index + 1
					} else {
						break
					}
				}
			}
		}
	}
	return heap[0]
}
