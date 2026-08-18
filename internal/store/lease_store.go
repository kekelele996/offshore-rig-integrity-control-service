package store

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"sync"
	"time"
)

type LeaseStore struct {
	mu     sync.Mutex
	leases map[string]domain.Lease
	now    func() time.Time
}

func NewLeaseStore(now func() time.Time) *LeaseStore {
	return &LeaseStore{leases: map[string]domain.Lease{}, now: now}
}
func (s *LeaseStore) Acquire(asset, owner string, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.leases[asset]; ok && old.ExpiresAt.After(s.now()) && old.Owner != owner {
		return domain.ErrLeaseConflict
	}
	s.leases[asset] = domain.Lease{AssetID: asset, Owner: owner, ExpiresAt: s.now().Add(ttl)}
	return nil
}
func (s *LeaseStore) Release(asset, owner string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.leases[asset]; ok && old.Owner == owner {
		delete(s.leases, asset)
	}
}
func (s *LeaseStore) Held(asset string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	old, ok := s.leases[asset]
	return ok && old.ExpiresAt.After(s.now())
}
