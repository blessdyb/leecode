package main

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	simulation := make([]int, 1001)
	for _, trip := range trips {
		for i := trip[1]; i < trip[2]; i++ {
			simulation[i] += trip[0]
		}
	}
	for _, s := range simulation {
		if s > capacity {
			return false
		}
	}
	return true
}

func carPoolingHeap(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[i][j]
	})
	heap := [][3]int{}
	down := func() {
		index := 0
		left := 2*index + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && heap[child+1][2] < heap[child][2] {
				child++
			}
			if heap[child][2] < heap[index][2] {
				heap[child], heap[index] = heap[index], heap[child]
				index = child
				left = 2*index + 1
			} else {
				break
			}
		}
	}
	space := capacity
	for i := 0; i < len(trips); i++ {
		for len(heap) > 0 && heap[0][2] <= trips[i][1] {
			space += heap[0][0]
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
			down()
		}
		if space < trips[i][0] {
			return false
		}
		heap = append(heap, [3]int{trips[i][0], trips[i][1], trips[i][2]})
		space -= trips[i][0]
		index := len(heap) - 1
		parent := (index - 1) / 2
		for index >= 0 && heap[index][2] < heap[parent][2] {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
			parent = (index - 1) / 2
		}
	}
	return true
}
