package main

import "sort"

func topKFrequentHeapSolution(words []string, k int) []string {
	hashmap := map[string]int{}
	for _, word := range words {
		hashmap[word]++
	}
	type WordFrequency struct {
		Word      string
		Frequency int
	}
	heap := []*WordFrequency{}
	isLarger := func(a, b *WordFrequency) bool {
		if a.Frequency > b.Frequency {
			return true
		} else if a.Frequency == b.Frequency {
			return a.Word < b.Word
		}
		return false
	}
	down := func() {
		parent := 0
		left := parent*2 + 1
		for left < len(heap) {
			child := left
			if child+1 < len(heap) && isLarger(heap[child], heap[child+1]) {
				child++
			}
			if isLarger(heap[parent], heap[child]) {
				heap[parent], heap[child] = heap[child], heap[parent]
				parent = child
				left = parent*2 + 1
			} else {
				break
			}
		}
	}
	for word, frequency := range hashmap {
		if len(heap) < k {
			heap = append(heap, &WordFrequency{
				Word:      word,
				Frequency: frequency,
			})
			index := len(heap) - 1
			parent := (index - 1) / 2
			for index >= 0 && isLarger(heap[parent], heap[index]) {
				heap[parent], heap[index] = heap[index], heap[parent]
				index = parent
				parent = (index - 1) / 2
			}
		} else if isLarger(&WordFrequency{
			Word:      word,
			Frequency: frequency,
		}, heap[0]) {
			heap[0] = &WordFrequency{
				Word:      word,
				Frequency: frequency,
			}
			down()
		}
	}
	ret := []string{}
	for len(heap) > 0 {
		ret = append([]string{heap[0].Word}, ret...)
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		down()
	}
	return ret
}

func topKFrequent(words []string, k int) []string {
	hashmap := map[string]int{}
	maxFrequency := 0
	for _, word := range words {
		hashmap[word]++
		if maxFrequency < hashmap[word] {
			maxFrequency = hashmap[word]
		}
	}
	bucket := make([][]string, maxFrequency+1)
	for key, frequency := range hashmap {
		bucket[frequency] = append(bucket[frequency], key)
	}
	ret := []string{}
	for maxFrequency > 0 && k > 0 {
		sort.Strings(bucket[maxFrequency])
		num := len(bucket[maxFrequency])
		if k < num {
			num = k
		}
		ret = append(ret, bucket[maxFrequency][:num]...)
		k -= num
		maxFrequency--
	}
	return ret
}
