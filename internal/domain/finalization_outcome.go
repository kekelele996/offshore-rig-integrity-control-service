package domain

// FinalizationOutcome reduces a primary finalization error and a cleanup error
// to the single error that callers should observe. The primary failure is
// authoritative: if it is non-nil, it is always surfaced. The cleanup error is
// only reported when the primary succeeded, so that a cleanup failure cannot
// mask the original finalization failure.
func FinalizationOutcome(primary, cleanup error) error {
	if primary != nil {
		return primary
	}
	return cleanup
}
