package worker

import (
	"context"
	"fmt"
)

func RunContextHandoff(ctx context.Context, work func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("handoff cancelled before dispatch: %w", err)
	}
	if err := work(ctx); err != nil {
		return fmt.Errorf("handoff work: %w", err)
	}
	return nil
}
