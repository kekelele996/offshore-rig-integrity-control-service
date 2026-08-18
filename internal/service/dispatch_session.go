package service

import "context"

func RunDispatchSession(ctx context.Context, zones []string) []string {
	child, cancel := context.WithCancel(ctx)
	results := make(chan string, len(zones))
	for _, zone := range zones {
		zone := zone
		go func() {
			select {
			case results <- zone:
			case <-child.Done():
			}
		}()
	}
	cancel()
	out := []string{}
	for i := 0; i < len(zones); i++ {
		select {
		case zone := <-results:
			out = append(out, zone)
		default:
			return out
		}
	}
	return out
}
