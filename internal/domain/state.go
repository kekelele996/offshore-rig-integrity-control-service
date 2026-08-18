package domain

var allowedTransitions = map[PlanState]map[PlanState]bool{
    PlanDraft: {PlanQueued:true, PlanRejected:true},
    PlanQueued: {PlanRunning:true, PlanRejected:true},
    PlanRunning: {PlanRejected:true, PlanReleased:true},
    PlanRejected: {},
    PlanReleased: {},
}

func CanTransition(from, to PlanState) bool { return allowedTransitions[from][to] }
func Transition(plan *InspectionPlan, to PlanState) error {
    if !CanTransition(plan.State, to) { return ErrInvalidState }
    plan.State=to; plan.Revision++; return nil
}
