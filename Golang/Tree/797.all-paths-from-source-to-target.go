package main

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := [][]int{}
	paths := [][]int{[]int{0}}
	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		current := path[len(path)-1]
		// It's a asyclic, so if we reached to the last node, no need to continue
		if current != len(graph)-1 {
			ret = append(ret, path)
		} else {
			for _, item := range graph[current] {
				paths = append(paths, append(append([]int{}, path...), item))
			}
		}
	}
	return ret
}
