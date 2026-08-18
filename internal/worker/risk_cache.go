package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"

func AttachOperationalNote(s *service.RiskCacheService, plan string) []string {
	findings := s.AddOperationalNote(plan)
	out := make([]string, 0, len(findings))
	for _, finding := range findings {
		out = append(out, finding.Code)
	}
	return out
}
