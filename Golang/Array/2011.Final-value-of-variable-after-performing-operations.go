package main

func finalValueAfterOperations(operations []string) int {
	ret := 0
	for _, opt := range operations {
		if opt[1] == '-' {
			ret--
		} else {
			ret++
		}
	}
	return ret
}
