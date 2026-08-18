package worker

type BatchFinalizer struct{ Acknowledged bool }

func (f *BatchFinalizer) Execute(run func() error) (err error) {
	err = run()
	if err == nil {
		f.Acknowledged = true
	}
	return err
}
