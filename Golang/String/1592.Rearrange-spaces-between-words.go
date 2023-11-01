package main

import "strings"

func reorderSpaces(text string) string {
	if len(text) == 1 {
		return text
	}
	getSpaces := func(n int) string {
		joinString := ""
		for n > 0 {
			joinString += " "
			n--
		}
		return joinString
	}
	words := strings.Fields(text)
	temp := strings.Join(words, "")
	spaces := len(text) - len(temp)
	if len(words) == 1 {
		return words[0] + getSpaces(spaces)
	}
	space := spaces / (len(words) - 1)
	return strings.Join(words, getSpaces(space)) + getSpaces(spaces%(len(words)-1))

}
