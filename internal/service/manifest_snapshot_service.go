package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func EmergencyManifestView(current domain.ManifestSnapshot, zone, contact string) domain.ManifestSnapshot {
	next := current
	next.Zones = append(append([]string(nil), current.Zones...), zone)
	next.Contacts = append(append([]string(nil), current.Contacts...), contact)
	next.RequiredChecks = make(map[string][]string, len(current.RequiredChecks)+1)
	for name, checks := range current.RequiredChecks {
		next.RequiredChecks[name] = append([]string(nil), checks...)
	}
	next.RequiredChecks[zone] = []string{"gas-test"}
	return next
}
