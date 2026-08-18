package worker

import (
 "context"; "sync"
 "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"
)
func Coordinate(ctx context.Context,jobs []domain.DispatchJob)([]string,error){in:=make(chan domain.DispatchJob);out:=make(chan string);var workers sync.WaitGroup;for i:=0;i<2;i++{workers.Add(1);go func(){defer workers.Done();for job:=range in{select{case out<-job.ID:case<-ctx.Done():return}}}()};go func(){for _,job:=range jobs{select{case in<-job:case<-ctx.Done():close(in);return}};close(in);workers.Wait();close(out)}();var done []string;for id:=range out{done=append(done,id)};if ctx.Err()!=nil{return nil,ctx.Err()};return done,nil}
