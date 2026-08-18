package service

import (
	"context"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

func (s *System) ApplyCancellation(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return domain.ErrCancelled
	default:
	}
	return s.Registry.Mutate(id, func(p *domain.InspectionPlan) error { return domain.Transition(p, domain.PlanQueued) })
}
