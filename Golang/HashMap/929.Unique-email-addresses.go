package main

func numUniqueEmails(emails []string) int {
	hashmap := map[string]bool{}
	for _, email := range emails {
		temp := []byte{}
		flag := false
		for i := 0; i < len(email); {
			if email[i] == '+' && !flag {
				for i < len(email) && email[i] != '@' {
					i++
				}
				continue
			} else if email[i] == '.' && !flag {
				i++
				continue
			} else if email[i] == '@' {
				flag = true
			}
			temp = append(temp, email[i])
			i++
		}
		hashmap[string(temp)] = true
	}
	return len(hashmap)
}
