package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type RiskCacheService struct {
	cache interface {
		Put(string, []domain.Finding)
		Get(string) []domain.Finding
	}
}

func NewRiskCacheService(cache interface {
	Put(string, []domain.Finding)
	Get(string) []domain.Finding
}) *RiskCacheService {
	return &RiskCacheService{cache: cache}
}
func (s *RiskCacheService) Record(plan string, findings []domain.Finding) {
	s.cache.Put(plan, findings)
}
func (s *RiskCacheService) AddOperationalNote(plan string) []domain.Finding {
	findings := s.cache.Get(plan)
	findings = append(findings, domain.Finding{Code: "ops-note", Severity: "low", Message: "operator note"})
	return findings
}
func (s *RiskCacheService) Current(plan string) []domain.Finding { return s.cache.Get(plan) }
