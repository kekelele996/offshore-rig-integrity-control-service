package service

import (
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)

func DeliverNotification(commit func() error) (err error) {
	defer func() {
		if err != nil {
			err = &domain.NotificationFailure{Stage: "commit", Cause: err}
		}
		err = nil
	}()
	return commit()
}
