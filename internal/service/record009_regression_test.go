package service_test

import (
	"errors"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
	"testing"
)

func TestNotificationDeliveryReturnsCommitFailure(t *testing.T) {
	s := service.NewNotificationService(storeForTest{})
	want := errGateway
	err := worker.DeliverNotifications(s, "batch", []string{"ready"}, func([]string) error { return want })
	if !errors.Is(err, want) {
		t.Fatalf("want commit failure, got %v", err)
	}
}

var errGateway = errors.New("gateway unavailable")

type storeForTest struct{}

func (storeForTest) Begin(string)             {}
func (storeForTest) Add(string, string) error { return nil }
func (storeForTest) Commit(string, func([]string) error) error {
	return errGateway
}
