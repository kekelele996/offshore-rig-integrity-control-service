package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type RiskProjector struct{ scratch []domain.Finding }

func (p *RiskProjector) Project(code string) []domain.Finding {
	p.scratch = append(p.scratch[:0], domain.Finding{Code: code, Severity: "low"})
	out := make([]domain.Finding, len(p.scratch))
	copy(out, p.scratch)
	return out
}
