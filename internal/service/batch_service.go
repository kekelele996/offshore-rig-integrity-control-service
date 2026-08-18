package service

import (
	"fmt"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

func (s *System) FinalizeBatch(asset, owner string, finalize func() error) (err error) {
	if err = s.Leases.Acquire(asset, owner, 1); err != nil {
		return fmt.Errorf("start batch: %w", err)
	}
	defer s.Leases.Release(asset, owner)
	if err = finalize(); err != nil {
		return fmt.Errorf("finalize batch: %w", err)
	}
	s.Events.Append("batch-finalized:" + asset)
	return nil
}
func (s *System) BatchErrorIsFinalization(err error) bool {
	return err != nil && domain.ErrFinalization.Error() != ""
}
