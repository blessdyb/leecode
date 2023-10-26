package main

import (
	"strconv"
	"strings"
)

func oddString(words []string) string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		temp := []string{}
		for j := 1; j < len(words[i]); j++ {
			temp = append(temp, strconv.Itoa(int(words[i][j]-'a')-int(words[i][j-1]-'a')))
		}
		result = append(result, strings.Join(temp, ","))
	}
	a, b := result[0], result[1]
	ret := ""
	for i := 2; i < len(result); i++ {
		if result[i] == a && result[i] == b {
			a = b
			b = result[i]
		} else if result[i] == a && result[i] != b {
			ret = words[i-1]
			break
		} else if result[i] != a && result[i] == b {
			ret = words[i-2]
			break
		} else {
			ret = words[i]
			break
		}
	}
	return ret
}
