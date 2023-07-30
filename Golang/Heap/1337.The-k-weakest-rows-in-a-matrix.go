package main

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	hashmap := map[int]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			hashmap[i] += mat[i][j]
		}
	}
	isWeaker := func(a, b [2]int) bool {
		if a[1] < b[1] {
			return true
		} else if a[1] == b[1] {
			return a[0] < b[0]
		}
		return false
	}
	heap := [][2]int{}
	down := func() {
		parent := 0
		left := parent*2 + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && isWeaker(heap[child], heap[child+1]) {
				child++
			}
			if isWeaker(heap[parent], heap[child]) {
				heap[child], heap[parent] = heap[parent], heap[child]
				parent = child
				left = parent*2 + 1
			} else {
				break
			}
		}
	}
	for key, val := range hashmap {
		if len(heap) < k {
			heap = append(heap, [2]int{key, val})
			index := len(heap) - 1
			parent := (index - 1) / 2
			for index >= 0 && isWeaker(heap[parent], heap[index]) {
				heap[index], heap[parent] = heap[parent], heap[index]
				index = parent
				parent = (index - 1) / 2
			}
		} else if isWeaker([2]int{key, val}, heap[0]) {
			heap[0] = [2]int{key, val}
			down()
		}
	}
	ret := []int{}
	for len(heap) > 0 {
		ret = append([]int{heap[0][0]}, ret...)
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		down()
	}
	return ret
}
