package service

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

type HandoverService struct {
	store interface {
		Save(string, []string)
		Read(string) []string
	}
}

func NewHandoverService(store interface {
	Save(string, []string)
	Read(string) []string
}) *HandoverService {
	return &HandoverService{store: store}
}
func (s *HandoverService) Start(plan string, recipients []string) {
	s.store.Save(plan, domain.CloneStrings(recipients))
}
func (s *HandoverService) AddContingencyRecipient(plan, recipient string) []string {
	recipients := s.store.Read(plan)
	recipients = append(recipients, recipient)
	return domain.CloneStrings(recipients)
}
func (s *HandoverService) StoredRecipients(plan string) []string { return s.store.Read(plan) }
