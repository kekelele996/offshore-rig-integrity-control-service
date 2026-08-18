package domain

func CloneStrings(values []string) []string {
	if values == nil {
		return nil
	}
	out := make([]string, len(values))
	copy(out, values)
	return out
}

func ClonePlan(in InspectionPlan) InspectionPlan {
	in.Zones = CloneStrings(in.Zones)
	return in
}

func HasZone(zones []string, wanted string) bool {
	for _, zone := range zones {
		if zone == wanted {
			return true
		}
	}
	return false
}
