package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func FanoutRisk(findings []domain.Finding) ([]domain.Finding, []domain.Finding) {
	left := make([]domain.Finding, len(findings))
	right := make([]domain.Finding, len(findings))
	copy(left, findings)
	copy(right, findings)
	return left, right
}
