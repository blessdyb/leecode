package main

func interpret(command string) string {
	ret := ""
	temp := ""
	for i := 0; i < len(command); i++ {
		temp += command[i : i+1]
		if temp == "G" {
			ret += "G"
			temp = ""
		} else if temp == "()" {
			ret += "o"
			temp = ""
		} else if temp == "(al)" {
			ret += "al"
			temp = ""
		}
	}
	return ret
}
