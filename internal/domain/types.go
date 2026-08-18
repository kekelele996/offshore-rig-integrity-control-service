package domain

import "time"

type PlanState string
const (
    PlanDraft PlanState = "draft"
    PlanQueued PlanState = "queued"
    PlanRunning PlanState = "running"
    PlanRejected PlanState = "rejected"
    PlanReleased PlanState = "released"
)

type InspectionPlan struct {
    ID string
    AssetID string
    Zones []string
    State PlanState
    Revision int
    CreatedAt time.Time
}

type DispatchJob struct { ID string; PlanID string; Zones []string; Attempt int }
type Lease struct { AssetID string; Owner string; ExpiresAt time.Time }
type Finding struct { Code string; Severity string; Message string }
type ReleaseDecision struct { PlanID string; Allowed bool; Reason string }

type PolicyProvider interface { Evaluate(InspectionPlan) ([]Finding, error) }
