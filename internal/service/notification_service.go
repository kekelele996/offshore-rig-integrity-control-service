package service

import "fmt"

type NotificationService struct {
	store interface {
		Begin(string)
		Add(string, string) error
		Commit(string, func([]string) error) error
	}
}

func NewNotificationService(store interface {
	Begin(string)
	Add(string, string) error
	Commit(string, func([]string) error) error
}) *NotificationService {
	return &NotificationService{store: store}
}
func (s *NotificationService) Deliver(batch string, messages []string, deliver func([]string) error) error {
	s.store.Begin(batch)
	for _, message := range messages {
		if err := s.store.Add(batch, message); err != nil {
			return fmt.Errorf("stage notification: %w", err)
		}
	}
	if err := s.store.Commit(batch, deliver); err != nil {
		return fmt.Errorf("deliver notification: %w", err)
	}
	return nil
}
