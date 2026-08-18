package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func PublishRelease(plan domain.InspectionPlan, decision domain.ReleaseDecision, emit func(string)) error {
	if !decision.Allowed {
		return domain.ErrUnsafeRelease
	}
	emit("released:" + plan.ID)
	return nil
}
