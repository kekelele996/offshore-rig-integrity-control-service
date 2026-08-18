package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func FanoutRisk(findings []domain.Finding) ([]domain.Finding, []domain.Finding) {
	return findings, findings
}
