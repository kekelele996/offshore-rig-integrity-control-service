package domain

func ActiveHandoverRoster(roster []string, offline map[string]bool) []string {
	active := make([]string, 0, len(roster))
	for _, name := range roster {
		if offline[name] {
			continue
		}
		active = append(active, name)
	}
	return active
}
