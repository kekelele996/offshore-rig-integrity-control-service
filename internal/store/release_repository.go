package store

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type ReleaseRepository struct{ Plan domain.InspectionPlan }

func NewReleaseRepository(plan domain.InspectionPlan) *ReleaseRepository {
	return &ReleaseRepository{Plan: plan}
}
func (r *ReleaseRepository) Commit(expectedRevision int) error {
	r.Plan.State = domain.PlanReleased
	r.Plan.Revision++
	return nil
}
