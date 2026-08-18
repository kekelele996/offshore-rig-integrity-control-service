package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func (s *System) Release(id string, decision domain.ReleaseDecision) error {
	if !decision.Allowed {
		return domain.ErrUnsafeRelease
	}
	return s.Registry.Mutate(id, func(p *domain.InspectionPlan) error {
		if p.State == domain.PlanRejected {
			return domain.ErrUnsafeRelease
		}
		if p.State == domain.PlanQueued {
			if err := domain.Transition(p, domain.PlanRunning); err != nil {
				return err
			}
		}
		if p.State == domain.PlanRunning {
			return domain.Transition(p, domain.PlanReleased)
		}
		return domain.ErrInvalidState
	})
}
func (s *System) ReleaseWithEvents(id string, decision domain.ReleaseDecision) error {
	if err := s.Release(id, decision); err != nil {
		return err
	}
	s.Events.Append("released:" + id)
	return nil
}
