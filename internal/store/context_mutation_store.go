package store

import (
	"context"
	"fmt"
	"sync"
)

type ContextMutationStore struct {
	mu        sync.Mutex
	committed []string
}

func NewContextMutationStore() *ContextMutationStore {
	return &ContextMutationStore{}
}

func (s *ContextMutationStore) Commit(ctx context.Context, planID string) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("commit plan %s: %w", planID, err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("commit plan %s: %w", planID, err)
	}
	s.committed = append(s.committed, planID)
	return nil
}

func (s *ContextMutationStore) Committed() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.committed)
}
