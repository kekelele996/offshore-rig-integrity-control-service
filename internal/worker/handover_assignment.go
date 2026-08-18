package worker

func AssignBackupCrew(roster []string, busy map[string]bool) []string {
	available := make([]string, 0, len(roster))
	for _, name := range roster {
		if busy[name] {
			continue
		}
		available = append(available, name)
	}
	return available
}
