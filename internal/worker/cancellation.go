package worker

import (
 "context"; "time"
 "github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
)
func CancelAwareMutation(ctx context.Context,s *service.System,id string)error{select{case<-ctx.Done():return ctx.Err();case<-time.After(time.Millisecond):};return s.ApplyCancellation(ctx,id)}
