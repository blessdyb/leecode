package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ret := 0
	for _, item := range items {
		if ruleKey == "type" && item[0] == ruleValue {
			ret++
		} else if ruleKey == "color" && item[1] == ruleValue {
			ret++
		} else if ruleKey == "name" && item[2] == ruleValue {
			ret++
		}
	}
	return ret
}
