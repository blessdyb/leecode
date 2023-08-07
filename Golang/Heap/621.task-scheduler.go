package main

func leastInterval(tasks []byte, n int) int {
	hashmap := map[byte]int{}
	for _, task := range tasks {
		hashmap[task]++
	}
	type CPUTask struct {
		Task  byte
		Count int
	}
	heap := []*CPUTask{}
	push := func(t *CPUTask) {
		heap = append(heap, t)
		index := len(heap) - 1
		parent := (index - 1) / 2
		for index >= 0 && heap[parent].Count < heap[index].Count {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
			parent = (index - 1) / 2
		}
	}
	pop := func() *CPUTask {
		temp := heap[0]
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		index := 0
		left := 2*index + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && heap[child+1].Count > heap[child].Count {
				child++
			}
			if heap[child].Count > heap[index].Count {
				heap[child], heap[index] = heap[index], heap[child]
				index = child
				left = 2*index + 1
			} else {
				break
			}
		}
		return temp
	}
	for k, v := range hashmap {
		push(&CPUTask{
			Task:  k,
			Count: v,
		})
	}
	ret := 0
	for len(heap) > 0 {
		stack := []*CPUTask{}
		cycle := n + 1
		for cycle > 0 && len(heap) > 0 {
			t := pop()
			if t.Count > 1 {
				t.Count--
				stack = append(stack, t)
			}
			ret++
			cycle--
		}
		if len(stack) > 0 {
			for _, s := range stack {
				push(s)
			}
			ret += cycle
		}
	}
	return ret
}
