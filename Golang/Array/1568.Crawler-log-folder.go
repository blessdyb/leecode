package main

func minOperations(logs []string) int {
	folders := []string{}
	for _, log := range logs {
		if log == "../" {
			if len(folders) > 0 {
				folders = folders[:len(folders)-1]
			}
		} else if log != "./" {
			folders = append(folders, log)
		}
	}
	return len(folders)
}
