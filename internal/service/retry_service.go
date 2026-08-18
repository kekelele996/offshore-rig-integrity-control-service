package service

import (
	"errors"
	"fmt"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

func (s *System) ReserveForRetry(asset, owner string) error {
	if err := s.Leases.Acquire(asset, owner, 1); err != nil {
		return fmt.Errorf("reserve inspection asset: %w", err)
	}
	return nil
}
func (s *System) RetryDisposition(err error) string {
	if errors.Is(err, domain.ErrLeaseConflict) {
		return "defer"
	}
	return "fail"
}
func (s *System) RetryAsset(asset, owner string) string {
	err := s.ReserveForRetry(asset, owner)
	return s.RetryDisposition(err)
}
