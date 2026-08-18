package handovercheck

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"reflect"
	"testing"
)

func TestHandoverContingencySnapshotRemainsDetached(t *testing.T) {
	want := []string{"marine", "electrical", "subsea", "medic"}
	domainInput := append([]string(nil), want...)
	_ = domain.ActiveHandoverRoster(domainInput, map[string]bool{"electrical": true})
	if !reflect.DeepEqual(domainInput, want) {
		t.Fatalf("domain compaction changed retained roster: %v", domainInput)
	}
	history := store.NewHandoverHistory(want)
	_ = history.Visible("electrical")
	if !reflect.DeepEqual(history.Roster(), want) {
		t.Fatalf("history compaction changed stored roster: %v", history.Roster())
	}
	serviceInput := append([]string(nil), want...)
	_ = service.ContingencyRoster(serviceInput, "marine")
	if !reflect.DeepEqual(serviceInput, want) {
		t.Fatalf("service compaction changed caller roster: %v", serviceInput)
	}
	workerInput := append([]string(nil), want...)
	_ = worker.AssignBackupCrew(workerInput, map[string]bool{"subsea": true})
	if !reflect.DeepEqual(workerInput, want) {
		t.Fatalf("worker compaction changed shared roster: %v", workerInput)
	}
}
