package service_test

import (
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

type nilSafePolicy struct{}

func (p *nilSafePolicy) Evaluate(domain.InspectionPlan) ([]domain.Finding, error) {
	if p == nil {
		return []domain.Finding{{Code: "nil-provider", Severity: "low"}}, nil
	}
	return nil, nil
}
func TestTypedNilPolicyRejectedAtEveryBoundary(t *testing.T) {
	var provider *nilSafePolicy
	var iface domain.PolicyProvider = provider
	if (domain.ProviderHandle{Provider: iface}).Available() {
		t.Fatal("domain handle reported a typed-nil provider as available")
	}
	registry := store.NewPolicyRegistry()
	if registry.Register("release", iface) || registry.Has("release") {
		t.Fatal("registry accepted a typed-nil provider")
	}
	if _, err := service.EvaluateProvider(domain.InspectionPlan{ID: "p4"}, iface); !errors.Is(err, domain.ErrPolicyUnavailable) {
		t.Fatalf("service evaluated typed-nil provider: %v", err)
	}
	if err := worker.AcceptPolicyJob(worker.PolicyJob{Provider: iface}); !errors.Is(err, domain.ErrPolicyUnavailable) {
		t.Fatalf("worker accepted typed-nil provider: %v", err)
	}
}
