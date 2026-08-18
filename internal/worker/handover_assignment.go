package worker

func AssignBackupCrew(roster []string, busy map[string]bool) []string {
	available := roster[:0]
	for index, name := range roster {
		if busy[name] {
			roster[index] = ""
			continue
		}
		available = append(available, name)
	}
	for index := len(available); index < len(roster); index++ {
		roster[index] = ""
	}
	return available
}
