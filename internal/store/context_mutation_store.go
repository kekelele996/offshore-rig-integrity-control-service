package store

import "context"

type ContextMutationStore struct{ committed []string }

func NewContextMutationStore() *ContextMutationStore { return &ContextMutationStore{} }
func (s *ContextMutationStore) Commit(ctx context.Context, planID string) error {
	s.committed = append(s.committed, planID)
	return nil
}
func (s *ContextMutationStore) Committed() int { return len(s.committed) }
