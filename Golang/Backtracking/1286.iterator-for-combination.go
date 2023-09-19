package main

type CombinationIterator struct {
	data []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combinations := []string{}
	combination := []byte{}
	var backtracking func(int)
	backtracking = func(i int) {
		if len(combination) == combinationLength {
			combinations = append(combinations, string(combination))
		}
		for j := i; j < len(characters); j++ {
			combination = append(combination, characters[j])
			backtracking(j + 1)
			combination = combination[:len(combination)-1]
		}
	}
	backtracking(0)
	return CombinationIterator{
		data: combinations,
	}
}

func (this *CombinationIterator) Next() string {
	ret := this.data[0]
	this.data = this.data[1:]
	return ret
}

func (this *CombinationIterator) HasNext() bool {
	return len(this.data) > 0
}
