package service

func FinalizeInspection(finalize, cleanup func() error) (err error) {
	defer func() { err = cleanup() }()
	err = finalize()
	return
}
