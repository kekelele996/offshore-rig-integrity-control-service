package service_test

import (
	"context"
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/assessment"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestCancelledContextStopsEveryMutationBoundary(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(errors.New("operator aborted"))
	if err := assessment.CheckAssessmentContext(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("assessment accepted cancelled context: %v", err)
	}
	direct := store.NewContextMutationStore()
	if err := direct.Commit(ctx, "p5-direct"); !errors.Is(err, context.Canceled) || direct.Committed() != 0 {
		t.Fatalf("store committed cancelled mutation: err=%v count=%d", err, direct.Committed())
	}
	throughService := store.NewContextMutationStore()
	if err := service.QueueWithContext(ctx, throughService, "p5-service"); !errors.Is(err, context.Canceled) || throughService.Committed() != 0 {
		t.Fatalf("service detached cancellation: err=%v count=%d", err, throughService.Committed())
	}
	called := false
	err := worker.RunContextHandoff(ctx, func(callCtx context.Context) error { called = true; return callCtx.Err() })
	if called || !errors.Is(err, context.Canceled) {
		t.Fatalf("worker ran cancelled handoff: called=%v err=%v", called, err)
	}
}
