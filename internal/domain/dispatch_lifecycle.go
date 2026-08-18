package domain

import (
	"context"
	"sync"
)

type CompletionGate struct {
	mu        sync.Mutex
	remaining int
	done      chan struct{}
	once      sync.Once
}

func NewCompletionGate(expected int) *CompletionGate {
	return &CompletionGate{remaining: expected, done: make(chan struct{})}
}
func (g *CompletionGate) Complete() {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.remaining > 0 {
		g.remaining--
	}
	if g.remaining == 0 {
		g.once.Do(func() { close(g.done) })
	}
}
func (g *CompletionGate) Wait(ctx context.Context) error {
	select {
	case <-g.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
