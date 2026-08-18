package domain

func HighestSeverity(findings []Finding) string {
	level := 0
	for _, finding := range findings {
		next := map[string]int{"low": 1, "medium": 2, "high": 3, "critical": 4}[finding.Severity]
		if next > level {
			level = next
		}
	}
	for name, score := range map[string]int{"low": 1, "medium": 2, "high": 3, "critical": 4} {
		if score == level {
			return name
		}
	}
	return "none"
}

func BlocksRelease(findings []Finding) bool {
	s := HighestSeverity(findings)
	return s == "high" || s == "critical"
}
