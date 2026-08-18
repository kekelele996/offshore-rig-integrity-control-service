package domain

type ReleaseAggregate struct {
	State    PlanState
	Revision int
}

func (a *ReleaseAggregate) ApplyRelease(expectedRevision int) error {
	if a.Revision != expectedRevision || a.State == PlanRejected || a.State == PlanReleased {
		return ErrInvalidState
	}
	if a.State != PlanRunning {
		return ErrUnsafeRelease
	}
	a.State = PlanReleased
	a.Revision++
	return nil
}
