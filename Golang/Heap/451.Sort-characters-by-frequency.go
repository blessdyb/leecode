package main

func frequencySort(s string) string {
	hashmap := map[rune]int{}
	for _, c := range s {
		hashmap[c]++
	}
	type PriorityQueue struct {
		Char   rune
		Weight int
	}
	heap := []PriorityQueue{}
	for key, value := range hashmap {
		heap = append(heap, PriorityQueue{
			Char:   key,
			Weight: value,
		})
		index := len(heap) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && heap[parentIndex].Weight < heap[index].Weight {
			heap[parentIndex], heap[index] = heap[index], heap[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	}
	ret := []rune{}
	for len(heap) > 0 {
		for i := 0; i < heap[0].Weight; i++ {
			ret = append(ret, heap[0].Char)
		}
		heap[0] = heap[len(heap)-1]
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < len(heap)-1 {
			child := left
			if child+1 < len(heap)-1 && heap[child+1].Weight > heap[child].Weight {
				child += 1
			}
			if heap[child].Weight > heap[parentIndex].Weight {
				heap[parentIndex], heap[child] = heap[child], heap[parentIndex]
				parentIndex = child
				left = parentIndex*2 + 1
			} else {
				break
			}
		}
		heap = heap[:len(heap)-1]
	}
	return string(ret)
}
