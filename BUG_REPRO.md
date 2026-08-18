# BUG_REPRO

## Bug 是什么

文件: internal/domain/dispatch_lifecycle.go、internal/store/dispatch_completion_store.go、internal/worker/zone_coordinator.go、internal/service/dispatch_session.go；符号: CompletionGate.Complete、DispatchCompletionStore.Acknowledge、CoordinateZones、RunDispatchSession；机制: 分布式派工在领域 gate、完成账本、worker 汇聚和服务 context 四处都过早宣布完成或取消，第一条结果会截断仍在运行的区位任务。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test -race ./zonelifecycle -run '^TestCoordinatorKeepsAllThreeInspectionZones$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestCoordinatorKeepsAllThreeInspectionZones (0.00s)
    record002_regression_test.go:46: domain completion gate opened before all zones finished
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/zonelifecycle	0.511s
FAIL
```
