package assessment

import "context"

func CheckAssessmentContext(ctx context.Context) error {
	if ctx == nil {
		return context.Canceled
	}
	return nil
}
