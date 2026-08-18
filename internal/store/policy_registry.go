package store

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"reflect"
)

type PolicyRegistry struct {
	providers map[string]domain.PolicyProvider
}

func NewPolicyRegistry() *PolicyRegistry {
	return &PolicyRegistry{providers: map[string]domain.PolicyProvider{}}
}
func providerPresent(p domain.PolicyProvider) bool {
	if p == nil {
		return false
	}
	value := reflect.ValueOf(p)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return !value.IsNil()
	default:
		return true
	}
}
func (r *PolicyRegistry) Register(name string, p domain.PolicyProvider) bool {
	if !providerPresent(p) {
		return false
	}
	r.providers[name] = p
	return true
}
func (r *PolicyRegistry) Has(name string) bool { return providerPresent(r.providers[name]) }
