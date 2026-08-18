package worker

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"reflect"
)

type PolicyJob struct {
	Plan     domain.InspectionPlan
	Provider domain.PolicyProvider
}

func AcceptPolicyJob(job PolicyJob) error {
	if job.Provider == nil {
		return domain.ErrPolicyUnavailable
	}
	value := reflect.ValueOf(job.Provider)
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return domain.ErrPolicyUnavailable
	}
	return nil
}
