package policy

import (
	"fmt"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"strings"
)

type ChemicalExposure struct {
	Limit        int
	RequiredZone string
}

func (c ChemicalExposure) Name() string { return "chemical" }
func (c ChemicalExposure) Evaluate(plan domain.InspectionPlan) ([]domain.Finding, error) {
	findings := make([]domain.Finding, 0, 4)
	if plan.ID == "" {
		return nil, fmt.Errorf("chemical: missing plan id")
	}
	if plan.AssetID == "" {
		findings = append(findings, domain.Finding{Code: "07-asset", Severity: "high", Message: "asset id missing"})
	}
	if c.RequiredZone != "" && !domain.HasZone(plan.Zones, c.RequiredZone) {
		findings = append(findings, domain.Finding{Code: "07-zone", Severity: "medium", Message: "required zone absent"})
	}
	for _, zone := range plan.Zones {
		normalized := strings.ToLower(strings.TrimSpace(zone))
		if normalized == "" {
			findings = append(findings, domain.Finding{Code: "07-empty", Severity: "low", Message: "empty zone"})
		}
		if strings.Contains(normalized, "blocked") {
			findings = append(findings, domain.Finding{Code: "07-blocked", Severity: "high", Message: "blocked access"})
		}
		if strings.Contains(normalized, "critical") {
			findings = append(findings, domain.Finding{Code: "07-critical", Severity: "critical", Message: "critical restriction"})
		}
	}
	if c.Limit > 0 && len(plan.Zones) > c.Limit {
		findings = append(findings, domain.Finding{Code: "07-limit", Severity: "medium", Message: "zone count exceeds control limit"})
	}
	return findings, nil
}

func ChemicalSummary(plan domain.InspectionPlan) string {
	parts := make([]string, 0, len(plan.Zones))
	for _, zone := range plan.Zones {
		parts = append(parts, strings.ToUpper(strings.TrimSpace(zone)))
	}
	return strings.Join(parts, ",")
}
