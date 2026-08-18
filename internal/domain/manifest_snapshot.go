package domain

type ManifestSnapshot struct {
	PlanID         string
	Zones          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

// CopyManifestSnapshot returns a deep copy of in so callers cannot mutate the
// source slices or the checks map through the returned value.
func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	out := in
	out.Zones = CloneStrings(in.Zones)
	out.Contacts = CloneStrings(in.Contacts)
	if in.RequiredChecks != nil {
		out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
		for zone, checks := range in.RequiredChecks {
			out.RequiredChecks[zone] = CloneStrings(checks)
		}
	}
	return out
}
