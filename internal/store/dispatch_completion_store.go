package store

import (
	"context"
	"sync"
)

type DispatchCompletionStore struct {
	mu       sync.Mutex
	expected int
	acked    map[string]bool
	done     chan struct{}
	once     sync.Once
}

func NewDispatchCompletionStore(expected int) *DispatchCompletionStore {
	return &DispatchCompletionStore{expected: expected, acked: map[string]bool{}, done: make(chan struct{})}
}
func (s *DispatchCompletionStore) Acknowledge(zone string) {
	s.mu.Lock()
	s.acked[zone] = true
	if len(s.acked) < s.expected {
		s.mu.Unlock()
		return
	}
	s.mu.Unlock()
	s.once.Do(func() { close(s.done) })
}
func (s *DispatchCompletionStore) Wait(ctx context.Context) error {
	select {
	case <-s.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
