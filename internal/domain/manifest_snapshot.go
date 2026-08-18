package domain

type ManifestSnapshot struct {
	PlanID         string
	Zones          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	return in
}
