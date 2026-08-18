package store

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"sync"
)

type HandoverStore struct {
	mu         sync.Mutex
	recipients map[string][]string
}

func NewHandoverStore() *HandoverStore { return &HandoverStore{recipients: map[string][]string{}} }
func (s *HandoverStore) Save(plan string, recipients []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.recipients[plan] = domain.CloneStrings(recipients)
}
func (s *HandoverStore) Read(plan string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return domain.CloneStrings(s.recipients[plan])
}
