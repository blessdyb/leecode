package main

func destCity(paths [][]string) string {
	hashmap := map[string]bool{}
	for _, path := range paths {
		hashmap[path[1]] = true
	}
	var ret string
	for _, path := range paths {
		if !hashmap[path[1]] {
			ret = path[1]
			break
		}
	}
	return ret
}
