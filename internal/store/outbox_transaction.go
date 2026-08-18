package store

type OutboxTransaction struct {
	RolledBack bool
	Committed  bool
}

func (t *OutboxTransaction) Finish(commit func() error) (err error) {
	defer func() {
		if err != nil {
			t.RolledBack = true
		}
	}()
	err = commit()
	if err == nil {
		t.Committed = true
	}
	return err
}
