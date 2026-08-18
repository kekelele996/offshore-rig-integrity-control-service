package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func copyDispatchManifest(in domain.ManifestSnapshot) domain.ManifestSnapshot {
	out := in
	out.Zones = append([]string(nil), in.Zones...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for zone, checks := range in.RequiredChecks {
		out.RequiredChecks[zone] = append([]string(nil), checks...)
	}
	return out
}

func DispatchManifest(snapshot domain.ManifestSnapshot) <-chan domain.ManifestSnapshot {
	out := make(chan domain.ManifestSnapshot, 1)
	owned := copyDispatchManifest(snapshot)
	go func() {
		out <- owned
		close(out)
	}()
	return out
}
