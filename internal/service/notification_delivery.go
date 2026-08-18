package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func DeliverNotification(commit func() error) (err error) {
	defer func() {
		if err != nil {
			err = &domain.NotificationFailure{
				Stage: "outbox commit",
				Cause: err,
			}
		}
	}()
	return commit()
}
