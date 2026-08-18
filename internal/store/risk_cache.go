package store

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"sync"
)

type RiskCache struct {
	mu       sync.RWMutex
	findings map[string][]domain.Finding
}

func NewRiskCache() *RiskCache { return &RiskCache{findings: map[string][]domain.Finding{}} }
func (c *RiskCache) Put(plan string, findings []domain.Finding) {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]domain.Finding, len(findings))
	copy(out, findings)
	c.findings[plan] = out
}
func (c *RiskCache) Get(plan string) []domain.Finding {
	c.mu.RLock()
	defer c.mu.RUnlock()
	values := c.findings[plan]
	out := make([]domain.Finding, len(values))
	copy(out, values)
	return out
}
