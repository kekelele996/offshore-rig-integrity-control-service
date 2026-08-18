package service_test

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"testing"
)

func TestHandoverContingencyRecipientDoesNotPolluteStoredList(t *testing.T) {
	s := service.NewHandoverService(store.NewHandoverStore())
	s.Start("p6", []string{"marine", "electrical"})
	got := s.AddContingencyRecipient("p6", "emergency")
	if len(got) != 3 {
		t.Fatal(got)
	}
	if stored := s.StoredRecipients("p6"); len(stored) != 2 {
		t.Fatalf("stored recipients polluted: %v", stored)
	}
}
