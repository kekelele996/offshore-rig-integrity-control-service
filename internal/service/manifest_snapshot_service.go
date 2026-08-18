package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

// EmergencyManifestView returns a new snapshot with the emergency zone appended
// without mutating the caller's snapshot, so the retained view stays stable.
func EmergencyManifestView(current domain.ManifestSnapshot, zone, contact string) domain.ManifestSnapshot {
	out := domain.CopyManifestSnapshot(current)
	out.Zones = append(out.Zones, zone)
	out.Contacts = append(out.Contacts, contact)
	out.RequiredChecks[zone] = append(out.RequiredChecks[zone], "gas-test")
	return out
}
