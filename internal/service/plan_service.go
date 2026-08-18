package service

import (
	"context"
	"fmt"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"time"
)

func (s *System) CreatePlan(id, asset string, zones []string) domain.InspectionPlan {
	p := domain.InspectionPlan{ID: id, AssetID: asset, Zones: domain.CloneStrings(zones), State: domain.PlanDraft, CreatedAt: s.now()}
	s.Registry.Put(p)
	return p
}
func (s *System) QueuePlan(ctx context.Context, id, owner string) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("queue: %v", err)
	}
	plan, ok := s.Registry.Get(id)
	if !ok {
		return domain.ErrInvalidState
	}
	if err := s.Leases.Acquire(plan.AssetID, owner, time.Minute); err != nil {
		return fmt.Errorf("queue lease: %v", err)
	}
	if err := s.Registry.Mutate(id, func(plan *domain.InspectionPlan) error { return domain.Transition(plan, domain.PlanQueued) }); err != nil {
		s.Leases.Release(plan.AssetID, owner)
		return err
	}
	s.Events.Append("queued:" + id)
	return nil
}
func (s *System) RejectPlan(id string) error {
	return s.Registry.Mutate(id, func(p *domain.InspectionPlan) error { return domain.Transition(p, domain.PlanRejected) })
}
