package service

import (
	"fmt"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"reflect"
)

func providerUnavailable(provider domain.PolicyProvider) bool {
	if provider == nil {
		return true
	}
	v := reflect.ValueOf(provider)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
func (s *System) EvaluateRelease(id string, provider domain.PolicyProvider) (domain.ReleaseDecision, error) {
	p, ok := s.Registry.Get(id)
	if !ok {
		return domain.ReleaseDecision{}, domain.ErrInvalidState
	}
	if providerUnavailable(provider) {
		return domain.ReleaseDecision{PlanID: id, Allowed: false, Reason: "policy unavailable"}, domain.ErrPolicyUnavailable
	}
	findings, err := provider.Evaluate(p)
	if err != nil {
		return domain.ReleaseDecision{}, fmt.Errorf("evaluate release: %w", err)
	}
	if domain.BlocksRelease(findings) {
		return domain.ReleaseDecision{PlanID: id, Allowed: false, Reason: "blocking finding"}, nil
	}
	return domain.ReleaseDecision{PlanID: id, Allowed: true}, nil
}
