package service

func ContingencyRoster(roster []string, primary string) []string {
	contingency := make([]string, 0, len(roster))
	for _, name := range roster {
		if name == primary {
			continue
		}
		contingency = append(contingency, name)
	}
	return contingency
}
