package service

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

func DeliverNotification(commit func() error) (err error) {
	err = commit()
	if err != nil {
		return &domain.NotificationFailure{Stage: "commit", Cause: err}
	}
	return nil
}
