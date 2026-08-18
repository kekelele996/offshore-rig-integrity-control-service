package worker

type BatchFinalizer struct{ Acknowledged bool }

func (f *BatchFinalizer) Execute(run func() error) (err error) {
	// A batch is acknowledged only when it completed without error. A failed
	// batch must not be acknowledged, otherwise upstream callers see success
	// and the failure never propagates to the next shift.
	defer func() {
		if err == nil {
			f.Acknowledged = true
		}
	}()
	return run()
}
