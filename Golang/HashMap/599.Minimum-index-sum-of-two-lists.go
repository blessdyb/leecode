package main

func findRestaurant(list1 []string, list2 []string) []string {
	h1 := map[string]int{}
	h2 := map[string]int{}
	for idx, str := range list1 {
		h1[str] += idx
		h2[str] |= 1
	}
	for idx, str := range list2 {
		h1[str] += idx
		h2[str] |= 2
	}
	max := 100000
	ret := []string{}
	for k, v := range h2 {
		if v == 3 {
			if h1[k] < max {
				ret = []string{k}
				max = h1[k]
			} else if h1[k] == max {
				ret = append(ret, k)
			}
		}
	}
	return ret
}
