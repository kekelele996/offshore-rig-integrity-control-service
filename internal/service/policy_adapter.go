package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func EvaluateProvider(plan domain.InspectionPlan, p domain.PolicyProvider) ([]domain.Finding, error) {
	if p == nil {
		return nil, domain.ErrPolicyUnavailable
	}
	return p.Evaluate(plan)
}
