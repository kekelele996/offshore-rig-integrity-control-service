package service

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
	"time"
)

type System struct {
	Registry  *store.Registry
	Leases    *store.LeaseStore
	Events    *store.EventStore
	Manifests *store.ManifestStore
	now       func() time.Time
}

func NewSystem(now func() time.Time) *System {
	return &System{Registry: store.NewRegistry(), Leases: store.NewLeaseStore(now), Events: store.NewEventStore(), Manifests: store.NewManifestStore(), now: now}
}
