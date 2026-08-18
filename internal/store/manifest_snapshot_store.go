package store

import (
	"sync"

	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

type ManifestSnapshotStore struct {
	mu    sync.Mutex
	items map[string]domain.ManifestSnapshot
}

func NewManifestSnapshotStore() *ManifestSnapshotStore {
	return &ManifestSnapshotStore{items: map[string]domain.ManifestSnapshot{}}
}

func copyStoredManifest(in domain.ManifestSnapshot) domain.ManifestSnapshot {
	out := in
	out.Zones = append([]string(nil), in.Zones...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for zone, checks := range in.RequiredChecks {
		out.RequiredChecks[zone] = append([]string(nil), checks...)
	}
	return out
}

func (s *ManifestSnapshotStore) Save(snapshot domain.ManifestSnapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[snapshot.PlanID] = copyStoredManifest(snapshot)
}

func (s *ManifestSnapshotStore) Load(planID string) domain.ManifestSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	return copyStoredManifest(s.items[planID])
}
