package service

// FinalizeInspection runs finalize, then always runs cleanup. The cleanup step
// must run on both the success and failure paths (the failure path is the one
// that was previously left incomplete). The returned error preserves the
// primary finalization failure: cleanup's error is surfaced only when
// finalization itself succeeded, so a successful cleanup can never overwrite a
// finalization failure with nil.
func FinalizeInspection(finalize, cleanup func() error) (err error) {
	err = finalize()
	cleanupErr := cleanup()
	if err == nil {
		err = cleanupErr
	}
	return err
}
