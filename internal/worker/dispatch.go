package worker

import (
	"context"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
)

func Dispatch(ctx context.Context, s *service.System, planID string) ([]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	zones, err := s.Registry.DispatchSnapshot(planID)
	if err != nil {
		return nil, err
	}
	jobs := make([]domain.DispatchJob, 0, len(zones))
	for _, zone := range zones {
		jobs = append(jobs, domain.DispatchJob{ID: planID + ":" + zone, PlanID: planID, Zones: []string{zone}})
	}
	return Coordinate(ctx, jobs)
}
