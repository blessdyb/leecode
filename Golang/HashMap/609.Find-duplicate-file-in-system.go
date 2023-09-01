package main

import "strings"

func findDuplicate(paths []string) [][]string {
	hashmap := map[string][]string{}
	for _, path := range paths {
		fnames := strings.Split(path, " ")
		for _, fname := range fnames[1:] {
			temp := strings.Split(fname, ".txt")
			content := temp[1]
			if _, ok := hashmap[content]; !ok {
				hashmap[content] = []string{}
			}
			hashmap[content] = append(hashmap[content], fnames[0]+"/"+temp[0]+".txt")
		}
	}
	ret := [][]string{}
	for _, v := range hashmap {
		if len(v) > 1 {
			ret = append(ret, v)
		}
	}
	return ret
}
