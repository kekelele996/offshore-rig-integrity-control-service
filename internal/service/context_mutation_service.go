package service

import (
	"context"
	"fmt"

	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
)

func QueueWithContext(ctx context.Context, mutations *store.ContextMutationStore, planID string) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("queue plan %s: %w", planID, err)
	}
	if err := mutations.Commit(ctx, planID); err != nil {
		return fmt.Errorf("queue plan %s: %w", planID, err)
	}
	return nil
}
