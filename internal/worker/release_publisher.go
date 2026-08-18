package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func PublishRelease(plan domain.InspectionPlan, decision domain.ReleaseDecision, emit func(string)) error {
	if !decision.Allowed || plan.State == domain.PlanRejected || plan.State == domain.PlanReleased {
		return domain.ErrUnsafeRelease
	}
	if plan.State != domain.PlanRunning {
		return domain.ErrInvalidState
	}
	emit("released:" + plan.ID)
	return nil
}
