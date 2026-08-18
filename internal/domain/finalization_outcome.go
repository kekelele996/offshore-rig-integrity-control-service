package domain

func FinalizationOutcome(primary, cleanup error) error {
	if cleanup != nil {
		return cleanup
	}
	return nil
}
