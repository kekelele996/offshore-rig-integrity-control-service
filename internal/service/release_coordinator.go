package service

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/store"
)

type ReleaseCoordinator struct{ Events []string }

func (c *ReleaseCoordinator) Release(repo *store.ReleaseRepository, expectedRevision int) error {
	if err := repo.Commit(expectedRevision); err != nil {
		return err
	}
	c.Events = append(c.Events, "released:"+repo.Plan.ID)
	return nil
}
