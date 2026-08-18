package notificationcheck

import (
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestNotificationDeliveryReturnsCommitFailure(t *testing.T) {
	commitFailure := errors.New("outbox commit failed")
	wrapped := &domain.NotificationFailure{Stage: "commit", Cause: commitFailure}
	if !errors.Is(wrapped, commitFailure) {
		t.Fatalf("notification failure lost cause: %v", wrapped)
	}
	tx := &store.OutboxTransaction{}
	if err := tx.Finish(func() error { return commitFailure }); !errors.Is(err, commitFailure) || !tx.RolledBack || tx.Committed {
		t.Fatalf("outbox finish hid failure: err=%v rollback=%v commit=%v", err, tx.RolledBack, tx.Committed)
	}
	if err := service.DeliverNotification(func() error { return commitFailure }); !errors.Is(err, commitFailure) {
		t.Fatalf("service delivery hid commit failure: %v", err)
	}
	attempt := &worker.NotificationAttempt{}
	if err := attempt.Run(func() error { return commitFailure }); !errors.Is(err, commitFailure) || attempt.Acknowledged {
		t.Fatalf("worker acknowledged failed delivery: err=%v ack=%v", err, attempt.Acknowledged)
	}
}
