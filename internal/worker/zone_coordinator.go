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
	out := make([]string, 0, len(zones))
	for len(out) < len(zones) {
		select {
		case zone := <-results:
			out = append(out, zone)
		case <-ctx.Done():
			return out
		}
	}
	return out
}
