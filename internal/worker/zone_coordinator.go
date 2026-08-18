package worker

import "context"

func CoordinateZones(ctx context.Context, zones []string) []string {
	results := make(chan string, len(zones))
	for _, zone := range zones {
		zone := zone
		go func() {
			select {
			case results <- zone:
			case <-ctx.Done():
			}
		}()
	}
	select {
	case first := <-results:
		return []string{first}
	case <-ctx.Done():
		return nil
	}
}
