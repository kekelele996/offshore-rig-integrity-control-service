package assessment

import (
	"context"
	"fmt"
)

func CheckAssessmentContext(ctx context.Context) error {
	if ctx == nil {
		return context.Canceled
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("assessment request stopped: %w", err)
	}
	return nil
}
