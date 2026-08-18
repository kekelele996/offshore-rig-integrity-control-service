package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func EmergencyManifestView(current domain.ManifestSnapshot, zone, contact string) domain.ManifestSnapshot {
	current.Zones = append(current.Zones, zone)
	current.Contacts = append(current.Contacts, contact)
	current.RequiredChecks[zone] = append(current.RequiredChecks[zone], "gas-test")
	return current
}
