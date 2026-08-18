package riskcachecheck

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestRiskCacheReadersReceiveIndependentSnapshots(t *testing.T) {
	window := &domain.FindingWindow{}
	first := window.Snapshot([]domain.Finding{{Code: "barrier", Severity: "low"}})
	_ = window.Snapshot([]domain.Finding{{Code: "weather", Severity: "high"}})
	if first[0].Code != "barrier" {
		t.Fatalf("domain window reused retained view: %+v", first)
	}
	cache := &store.RiskCache{}
	cache.Put([]domain.Finding{{Code: "corrosion", Severity: "medium"}})
	cached := cache.Read()
	cache.Put([]domain.Finding{{Code: "fatigue", Severity: "high"}})
	if cached[0].Code != "corrosion" {
		t.Fatalf("cache read was overwritten: %+v", cached)
	}
	projector := &service.RiskProjector{}
	projected := projector.Project("lifting")
	_ = projector.Project("subsea")
	if projected[0].Code != "lifting" {
		t.Fatalf("projector reused consumer view: %+v", projected)
	}
	source := []domain.Finding{{Code: "permit", Severity: "low"}}
	left, right := worker.FanoutRisk(source)
	right[0].Code = "consumer-overwrite"
	if left[0].Code != "permit" {
		t.Fatalf("fanout consumers shared a live buffer: %+v", left)
	}
}
