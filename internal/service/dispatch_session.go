package service

import (
	"context"
	"sync"
)

func RunDispatchSession(ctx context.Context, zones []string) []string {
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan string, len(zones))
	var workers sync.WaitGroup
	workers.Add(len(zones))
	for _, zone := range zones {
		zone := zone
		go func() {
			defer workers.Done()
			select {
			case results <- zone:
			case <-child.Done():
			}
		}()
	}
	go func() { workers.Wait(); close(results) }()
	out := make([]string, 0, len(zones))
	for zone := range results {
		out = append(out, zone)
	}
	return out
}
