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

// Save stores a deep copy so later mutations to the caller's snapshot cannot
// alter the stored record.
func (s *ManifestSnapshotStore) Save(snapshot domain.ManifestSnapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[snapshot.PlanID] = domain.CopyManifestSnapshot(snapshot)
}

// Load returns a deep copy so callers cannot mutate the stored record through
// the returned value.
func (s *ManifestSnapshotStore) Load(planID string) domain.ManifestSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	return domain.CopyManifestSnapshot(s.items[planID])
}
