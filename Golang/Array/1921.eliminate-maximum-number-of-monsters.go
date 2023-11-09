package main

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	ret := 0
	init := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		init[i] = float64(dist[i]) / float64(speed[i])
	}
	sort.Float64s(init)
	for i := 0; i < len(init); i++ {
		if init[i] <= float64(i) {
			break
		}
		ret++
	}
	return ret
}

func eliminateMaximumSpeedHeap(dist []int, speed []int) int {
	ret := 0
	init := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		init[i] = float64(dist[i]) / float64(speed[i])
	}
	heap := []float64{}
	for i := 0; i < len(init); i++ {
		heap = append(heap, init[i])
		index := len(heap) - 1
		parent := (index - 1) / 2
		for index > 0 && heap[parent] > heap[index] {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
		}
	}
	for {
		if len(heap) > 0 && heap[0] > float64(ret) {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
			index := 0
			child := 2*index + 1
			for child < len(heap) {
				right := child + 1
				if right < len(heap) && heap[right] < heap[child] {
					child = right
				}
				if heap[child] < heap[index] {
					heap[child], heap[index] = heap[index], heap[child]
					index = child
					child = 2*index + 1
				} else {
					break
				}
			}
			ret++
		} else {
			break
		}
	}
	return ret
}
