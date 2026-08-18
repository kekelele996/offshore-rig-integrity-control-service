package service

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"reflect"
)

func EvaluateProvider(plan domain.InspectionPlan, p domain.PolicyProvider) ([]domain.Finding, error) {
	if p == nil {
		return nil, domain.ErrPolicyUnavailable
	}
	value := reflect.ValueOf(p)
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return nil, domain.ErrPolicyUnavailable
	}
	return p.Evaluate(plan)
}
