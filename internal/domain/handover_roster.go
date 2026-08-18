package domain

func ActiveHandoverRoster(roster []string, offline map[string]bool) []string {
	active := roster[:0]
	for index := range roster {
		name := roster[index]
		if offline[name] {
			roster[index] = ""
			continue
		}
		active = append(active, name)
	}
	for index := len(active); index < len(roster); index++ {
		roster[index] = ""
	}
	return active
}
