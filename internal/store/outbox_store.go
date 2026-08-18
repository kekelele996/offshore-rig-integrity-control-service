package store

import (
	"fmt"
	"sync"
)

type OutboxStore struct {
	mu       sync.Mutex
	open     map[string]bool
	messages map[string][]string
}

func NewOutboxStore() *OutboxStore {
	return &OutboxStore{open: map[string]bool{}, messages: map[string][]string{}}
}
func (s *OutboxStore) Begin(batch string) { s.mu.Lock(); defer s.mu.Unlock(); s.open[batch] = true }
func (s *OutboxStore) Add(batch, message string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.open[batch] {
		return fmt.Errorf("outbox %s is closed", batch)
	}
	s.messages[batch] = append(s.messages[batch], message)
	return nil
}
func (s *OutboxStore) Commit(batch string, deliver func([]string) error) error {
	s.mu.Lock()
	messages := append([]string(nil), s.messages[batch]...)
	open := s.open[batch]
	s.mu.Unlock()
	if !open {
		return fmt.Errorf("outbox %s is closed", batch)
	}
	if err := deliver(messages); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.open[batch] = false
	return nil
}
func (s *OutboxStore) IsOpen(batch string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.open[batch]
}
