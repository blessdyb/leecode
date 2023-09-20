package main

func thirdMax(nums []int) int {
	heap := []int{}
	hashmap := map[int]bool{}
	length := 3
	for _, num := range nums {
		if !hashmap[num] {
			hashmap[num] = true
			if len(heap) < length {
				heap = append(heap, num)
				index := len(heap) - 1
				parent := (index - 1) / 2
				for index >= 0 && heap[index] < heap[parent] {
					heap[index], heap[parent] = heap[parent], heap[index]
					index = parent
					parent = (index - 1) / 2
				}
			} else if heap[0] < num {
				heap[0] = num
				index := 0
				left := 2*index + 1
				for left < length {
					child := left
					if child+1 < length && heap[child] > heap[child+1] {
						child++
					}
					if heap[index] > heap[child] {
						heap[child], heap[index] = heap[index], heap[child]
						index = child
						left = 2*index + 1
					} else {
						break
					}
				}
			}
		}
	}
	if len(heap) == length {
		return heap[0]
	}
	return heap[len(heap)-1]
}
