package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	hashmap := map[string]int{}
	for _, cpdomain := range cpdomains {
		list := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(list[0])
		dList := strings.Split(list[1], ".")
		d := dList[len(dList)-1]
		hashmap[d] += count
		for i := len(dList) - 2; i >= 0; i-- {
			d = dList[i] + "." + d
			hashmap[d] += count
		}
	}
	ret := []string{}
	for k, v := range hashmap {
		ret = append(ret, fmt.Sprintf("%d %s", v, k))
	}
	return ret
}
