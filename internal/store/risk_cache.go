package store

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type RiskCache struct {
	findings []domain.Finding
}

func (c *RiskCache) Put(findings []domain.Finding) {
	c.findings = append(c.findings[:0], findings...)
}

func (c *RiskCache) Read() []domain.Finding {
	snapshot := make([]domain.Finding, len(c.findings))
	copy(snapshot, c.findings)
	return snapshot
}
