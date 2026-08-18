package domain

type ReleaseAggregate struct {
	State    PlanState
	Revision int
}

func (a *ReleaseAggregate) ApplyRelease(expectedRevision int) error {
	a.State = PlanReleased
	a.Revision++
	return nil
}
