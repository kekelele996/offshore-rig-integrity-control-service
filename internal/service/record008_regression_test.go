package service_test

import (
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestRejectedReleaseIsFencedAtEveryBoundary(t *testing.T) {
	aggregate := domain.ReleaseAggregate{State: domain.PlanRejected, Revision: 7}
	if err := aggregate.ApplyRelease(7); !errors.Is(err, domain.ErrInvalidState) || aggregate.State != domain.PlanRejected {
		t.Fatalf("aggregate released terminal state: %+v err=%v", aggregate, err)
	}
	plan := domain.InspectionPlan{ID: "p8", State: domain.PlanRejected, Revision: 7}
	repo := store.NewReleaseRepository(plan)
	if err := repo.Commit(7); !errors.Is(err, domain.ErrInvalidState) || repo.Plan.State != domain.PlanRejected {
		t.Fatalf("repository released rejected plan: %+v err=%v", repo.Plan, err)
	}
	repo2 := store.NewReleaseRepository(plan)
	coordinator := &service.ReleaseCoordinator{}
	if err := coordinator.Release(repo2, 7); !errors.Is(err, domain.ErrInvalidState) || len(coordinator.Events) != 0 {
		t.Fatalf("coordinator emitted before rejected commit: events=%v err=%v", coordinator.Events, err)
	}
	published := []string{}
	err := worker.PublishRelease(plan, domain.ReleaseDecision{PlanID: "p8", Allowed: true}, func(event string) { published = append(published, event) })
	if !errors.Is(err, domain.ErrUnsafeRelease) || len(published) != 0 {
		t.Fatalf("publisher emitted rejected release: events=%v err=%v", published, err)
	}
}
