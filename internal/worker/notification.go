package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"

func DeliverNotifications(s *service.NotificationService, batch string, messages []string, send func([]string) error) error {
	return s.Deliver(batch, messages, send)
}
