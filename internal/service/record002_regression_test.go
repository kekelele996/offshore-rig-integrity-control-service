package service_test

import (
	"context"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
	"time"
)

func TestCoordinatorKeepsAllThreeInspectionZones(t *testing.T) {
	s := fresh()
	s.CreatePlan("three", "asset", []string{"north", "south", "west"})
	if err := s.QueuePlan(context.Background(), "three", "operator"); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	got, err := worker.Dispatch(ctx, s, "three")
	if err != nil || len(got) != 3 {
		t.Fatalf("jobs=%v err=%v", got, err)
	}
}
