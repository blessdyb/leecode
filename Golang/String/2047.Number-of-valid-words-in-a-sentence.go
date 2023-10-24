package main

import (
	"fmt"
	"strings"
)

func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	count := 0
	for i := 0; i < len(words); i++ {
		flag := true
		punctuation_hyphen := 0
		punctuation_rest := 0
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] >= '0' && words[i][j] <= '9' {
				flag = false
				break
			} else if words[i][j] == '-' {
				if j == 0 || j == len(words[i])-1 || (words[i][j-1] < 'a' || words[i][j-1] > 'z') || (words[i][j+1] < 'a' || words[i][j+1] > 'z') {
					flag = false
					break
				} else {
					punctuation_hyphen++
				}
			} else if words[i][j] == ',' || words[i][j] == '!' || words[i][j] == '.' {

				if j != len(words[i])-1 {
					flag = false
					break
				} else {
					punctuation_rest++
				}
			}
		}
		if flag && punctuation_hyphen < 2 && punctuation_rest < 2 {
			count++
			fmt.Println(words[i])
		}
	}
	return count
}
