package store

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type PolicyRegistry struct {
	providers map[string]domain.PolicyProvider
}

func NewPolicyRegistry() *PolicyRegistry {
	return &PolicyRegistry{providers: map[string]domain.PolicyProvider{}}
}
func (r *PolicyRegistry) Register(name string, p domain.PolicyProvider) bool {
	if p == nil {
		return false
	}
	r.providers[name] = p
	return true
}
func (r *PolicyRegistry) Has(name string) bool { return r.providers[name] != nil }
