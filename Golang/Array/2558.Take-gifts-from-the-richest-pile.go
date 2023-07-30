package main

import (
	"fmt"
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	lastIndex := len(gifts) - 1
	for k > 0 {
		sort.Ints(gifts)
		gifts[lastIndex] = int(math.Sqrt(float64(gifts[lastIndex])))
		k--
	}
	fmt.Println(gifts)
	sum := 0
	for _, g := range gifts {
		sum += g
	}
	return int64(sum)
}
