package store

type OutboxTransaction struct {
	RolledBack bool
	Committed  bool
}

func (t *OutboxTransaction) Finish(commit func() error) (err error) {
	err = commit()
	if err != nil {
		t.RolledBack = true
		return err
	}
	t.Committed = true
	return nil
}
