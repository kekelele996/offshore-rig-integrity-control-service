package service

import (
	"context"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
)

func QueueWithContext(ctx context.Context, s *store.ContextMutationStore, planID string) error {
	return s.Commit(context.WithoutCancel(ctx), planID)
}
