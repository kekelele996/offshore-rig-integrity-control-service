package finalizationcheck

import (
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestBatchFinalizationPreservesFailureAcrossCleanup(t *testing.T) {
	primary := domain.ErrFinalization
	if err := domain.FinalizationOutcome(primary, nil); !errors.Is(err, primary) {
		t.Fatalf("domain outcome lost primary failure: %v", err)
	}
	lease := store.NewFinalizationLease()
	if err := lease.Run(func() error { return primary }); !errors.Is(err, primary) || lease.Held() {
		t.Fatalf("lease cleanup failed: err=%v held=%v", err, lease.Held())
	}
	if err := service.FinalizeInspection(func() error { return primary }, func() error { return nil }); !errors.Is(err, primary) {
		t.Fatalf("service cleanup overwrote failure: %v", err)
	}
	finalizer := &worker.BatchFinalizer{}
	if err := finalizer.Execute(func() error { return primary }); !errors.Is(err, primary) || finalizer.Acknowledged {
		t.Fatalf("worker acknowledged failed batch: err=%v ack=%v", err, finalizer.Acknowledged)
	}
}
