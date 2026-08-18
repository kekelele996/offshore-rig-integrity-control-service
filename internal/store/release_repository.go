package store

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type ReleaseRepository struct{ Plan domain.InspectionPlan }

func NewReleaseRepository(plan domain.InspectionPlan) *ReleaseRepository {
	return &ReleaseRepository{Plan: plan}
}
func (r *ReleaseRepository) Commit(expectedRevision int) error {
	if r.Plan.Revision != expectedRevision || r.Plan.State == domain.PlanRejected || r.Plan.State == domain.PlanReleased {
		return domain.ErrInvalidState
	}
	if r.Plan.State != domain.PlanRunning {
		return domain.ErrUnsafeRelease
	}
	r.Plan.State = domain.PlanReleased
	r.Plan.Revision++
	return nil
}
