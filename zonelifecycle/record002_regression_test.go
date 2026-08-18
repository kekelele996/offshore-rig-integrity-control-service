package workercheck

import (
	"context"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
)

func TestCoordinatorKeepsAllThreeInspectionZones(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	gate := domain.NewCompletionGate(3)
	releaseFollowers := make(chan struct{})
	firstCompleted := make(chan struct{})
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(releaseFollowers) }) }
	defer release()
	var participants sync.WaitGroup
	participants.Add(3)
	go func() {
		defer participants.Done()
		gate.Complete()
		close(firstCompleted)
	}()
	go func() {
		defer participants.Done()
		<-releaseFollowers
		gate.Complete()
	}()
	go func() {
		defer participants.Done()
		<-releaseFollowers
		gate.Complete()
	}()
	<-firstCompleted
	short, cancelShort := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelShort()
	if gate.Wait(short) == nil {
		t.Fatal("domain completion gate opened before all zones finished")
	}
	release()
	participants.Wait()
	if err := gate.Wait(ctx); err != nil {
		t.Fatalf("domain completion gate never opened: %v", err)
	}
	ledger := store.NewDispatchCompletionStore(3)
	ledger.Acknowledge("north")
	short2, cancelShort2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelShort2()
	if ledger.Wait(short2) == nil {
		t.Fatal("dispatch ledger completed after one acknowledgement")
	}
	ledger.Acknowledge("south")
	ledger.Acknowledge("subsea")
	if err := ledger.Wait(ctx); err != nil {
		t.Fatalf("dispatch ledger never completed: %v", err)
	}
	want := []string{"north", "south", "subsea"}
	got := worker.CoordinateZones(ctx, want)
	sort.Strings(got)
	if len(got) != 3 {
		t.Fatalf("worker coordinator returned %v", got)
	}
	session := service.RunDispatchSession(ctx, want)
	sort.Strings(session)
	if len(session) != 3 {
		t.Fatalf("service session returned %v", session)
	}
}
