package domain

type ManifestSnapshot struct {
	PlanID         string
	Zones          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	out := in
	out.Zones = append([]string(nil), in.Zones...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for zone, checks := range in.RequiredChecks {
		out.RequiredChecks[zone] = append([]string(nil), checks...)
	}
	return out
}
