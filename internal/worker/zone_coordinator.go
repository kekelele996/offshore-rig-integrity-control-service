package worker

import (
	"context"
	"sync"
)

func CoordinateZones(ctx context.Context, zones []string) []string {
	results := make(chan string, len(zones))
	var workers sync.WaitGroup
	for _, zone := range zones {
		zone := zone
		workers.Add(1)
		go func() {
			defer workers.Done()
			select {
			case results <- zone:
			case <-ctx.Done():
			}
		}()
	}
	go func() {
		workers.Wait()
		close(results)
	}()
	out := make([]string, 0, len(zones))
	for zone := range results {
		out = append(out, zone)
	}
	return out
}
