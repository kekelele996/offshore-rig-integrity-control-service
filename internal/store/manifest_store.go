package store

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"sync"
)

type ManifestStore struct {
	mu    sync.Mutex
	zones map[string][]string
}

func NewManifestStore() *ManifestStore { return &ManifestStore{zones: map[string][]string{}} }
func (s *ManifestStore) Save(plan string, zones []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.zones[plan] = domain.CloneStrings(zones)
}
func (s *ManifestStore) Read(plan string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return domain.CloneStrings(s.zones[plan])
}
