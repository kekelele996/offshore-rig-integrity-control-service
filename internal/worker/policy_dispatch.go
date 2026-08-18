package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type PolicyJob struct {
	Plan     domain.InspectionPlan
	Provider domain.PolicyProvider
}

func AcceptPolicyJob(job PolicyJob) error {
	if job.Provider == nil {
		return domain.ErrPolicyUnavailable
	}
	return nil
}
