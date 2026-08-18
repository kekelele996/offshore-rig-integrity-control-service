package worker

type BatchFinalizer struct{ Acknowledged bool }

func (f *BatchFinalizer) Execute(run func() error) (err error) {
	defer func() { f.Acknowledged = true }()
	return run()
}
