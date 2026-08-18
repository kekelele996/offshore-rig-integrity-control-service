package worker

import "context"

func RunContextHandoff(ctx context.Context, work func(context.Context) error) error {
	return work(context.Background())
}
