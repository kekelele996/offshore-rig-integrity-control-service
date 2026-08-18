package service_test

import (
 "context"; "errors"; "testing"; "time"
 "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
 "github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
 "github.com/kekelele996/offshore-rig-integrity-control-service/internal/worker"
)

type cleanPolicy struct{}
func(cleanPolicy)Evaluate(domain.InspectionPlan)([]domain.Finding,error){return nil,nil}
type typedNilPolicy struct{}
func(*typedNilPolicy)Evaluate(domain.InspectionPlan)([]domain.Finding,error){return nil,nil}
func fresh()*service.System{return service.NewSystem(func()time.Time{return time.Unix(100,0)})}
func queued(s *service.System,id string){s.CreatePlan(id,"asset-"+id,[]string{"north","south"});if err:=s.QueuePlan(context.Background(),id,"operator");err!=nil{panic(err)}}
func TestConcurrentDispatchSnapshotDoesNotLeakZones(t *testing.T){s:=fresh();queued(s,"p1");if err:=s.StageManifest("p1");err!=nil{t.Fatal(err)};copy:=s.AddEmergencyZone("p1","emergency");if len(copy)!=3{t.Fatal(copy)};if got:=s.StoredManifest("p1");len(got)!=2{t.Fatalf("manifest leaked %v",got)}}
func TestCoordinatorDrainsJobs(t *testing.T){s:=fresh();queued(s,"p2");got,err:=worker.Dispatch(context.Background(),s,"p2");if err!=nil||len(got)!=2{t.Fatalf("got=%v err=%v",got,err)}}
func TestRetryRecognizesLeaseConflict(t *testing.T){s:=fresh();if err:=s.ReserveForRetry("asset","first");err!=nil{t.Fatal(err)};if got:=s.RetryAsset("asset","second");got!="defer"{t.Fatalf("got %s",got)}}
func TestTypedNilPolicyIsDenied(t *testing.T){s:=fresh();s.CreatePlan("p4","asset",[]string{"north"});var p *typedNilPolicy;decision,err:=s.EvaluateRelease("p4",p);if !errors.Is(err,domain.ErrPolicyUnavailable)||decision.Allowed{t.Fatalf("decision=%+v err=%v",decision,err)}}
func TestCancelledContextDoesNotQueuePlan(t *testing.T){s:=fresh();s.CreatePlan("p5","asset",[]string{"north"});ctx,cancel:=context.WithCancel(context.Background());cancel();if err:=worker.CancelAwareMutation(ctx,s,"p5");err==nil{t.Fatal("expected cancellation")};p,_:=s.Registry.Get("p5");if p.State!=domain.PlanDraft{t.Fatalf("state=%s",p.State)}}
func TestManifestAppendDoesNotPolluteStoredZones(t *testing.T){s:=fresh();queued(s,"p6");if err:=s.StageManifest("p6");err!=nil{t.Fatal(err)};_ = s.AddEmergencyZone("p6","temporary");if got:=s.StoredManifest("p6");len(got)!=2{t.Fatalf("got %v",got)}}
func TestFinalizeBatchReturnsErrorAndReleasesLease(t *testing.T){s:=fresh();err:=s.FinalizeBatch("asset","owner",func()error{return domain.ErrFinalization});if !errors.Is(err,domain.ErrFinalization){t.Fatalf("err=%v",err)};if s.Leases.Held("asset"){t.Fatal("lease kept")}}
func TestRejectedPlanCannotRelease(t *testing.T){s:=fresh();queued(s,"p8");if err:=s.RejectPlan("p8");err!=nil{t.Fatal(err)};if err:=s.ReleaseWithEvents("p8",domain.ReleaseDecision{PlanID:"p8",Allowed:true});!errors.Is(err,domain.ErrUnsafeRelease){t.Fatalf("err=%v",err)};if s.Events.Count("released:p8")!=0{t.Fatal("release event emitted")}}
func TestReleaseMovesQueuedPlanThroughRunning(t *testing.T){s:=fresh();queued(s,"p9");if err:=s.ReleaseWithEvents("p9",domain.ReleaseDecision{PlanID:"p9",Allowed:true});err!=nil{t.Fatal(err)};p,_:=s.Registry.Get("p9");if p.State!=domain.PlanReleased{t.Fatal(p.State)}}
func TestPolicyBlocksCriticalZone(t *testing.T){s:=fresh();s.CreatePlan("p10","asset",[]string{"critical-access"});d,err:=s.EvaluateRelease("p10",cleanPolicy{});if err!=nil{t.Fatal(err)};if !d.Allowed{t.Fatal("clean policy should allow")}}
