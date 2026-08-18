# BUG_REPRO

## Bug 是什么

文件: internal/assessment/context_barrier.go、internal/store/context_mutation_store.go、internal/service/context_mutation_service.go、internal/worker/context_handoff.go；符号: CheckAssessmentContext、ContextMutationStore.Commit、QueueWithContext、RunContextHandoff；机制: 评估、仓储、服务与 worker 边界分别忽略 ctx.Err、使用 WithoutCancel 或替换为 Background，取消链在跨层传递中断裂并允许状态写入。

## 如何触发

在项目根目录执行下面这条定向回归命令：

```bash
go test ./internal/service -run '^TestCancelledContextStopsEveryMutationBoundary$' -count=1
```

## 错误信息

埋错环境的真实输出如下：

```text
--- FAIL: TestCancelledContextStopsEveryMutationBoundary (0.00s)
    record005_regression_test.go:17: assessment accepted cancelled context: <nil>
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/internal/service	0.345s
FAIL
```
