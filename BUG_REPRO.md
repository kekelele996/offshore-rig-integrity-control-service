# BUG_REPRO

## Bug 是什么

文件: internal/domain/release_aggregate.go、internal/store/release_repository.go、internal/service/release_coordinator.go、internal/worker/release_publisher.go；符号: ReleaseAggregate.ApplyRelease、ReleaseRepository.Commit、ReleaseCoordinator.Release、PublishRelease；机制: rejected 终态与 revision 栅栏在聚合、仓储提交、服务事件顺序和发布器四处均未执行，导致终态计划被写成 released 并提前发出事件。

## 如何触发

在项目根目录执行下面这条定向回归命令：

```bash
go test ./internal/service -run '^TestRejectedReleaseIsFencedAtEveryBoundary$' -count=1
```

## 错误信息

埋错环境的真实输出如下：

```text
--- FAIL: TestRejectedReleaseIsFencedAtEveryBoundary (0.00s)
    record008_regression_test.go:15: aggregate released terminal state: {State:released Revision:8} err=<nil>
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/internal/service	0.318s
FAIL
```
