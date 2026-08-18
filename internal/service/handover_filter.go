package service

func ContingencyRoster(roster []string, primary string) []string {
	contingency := roster[:0]
	for index, name := range roster {
		if name == primary {
			roster[index] = ""
			continue
		}
		contingency = append(contingency, name)
	}
	for index := len(contingency); index < len(roster); index++ {
		roster[index] = ""
	}
	return contingency
}
