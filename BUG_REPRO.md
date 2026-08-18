# BUG_REPRO

## Bug 是什么

文件: internal/domain/finalization_outcome.go、internal/store/finalization_lease.go、internal/service/finalization_pipeline.go、internal/worker/batch_finalizer.go；符号: FinalizationOutcome、FinalizationLease.Run、FinalizeInspection、BatchFinalizer.Execute；机制: defer 清理链在领域、仓储、服务和 worker 四层分别丢弃主错误、仅成功时释放、覆盖命名返回值并错误确认失败批次。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test ./batchcleanup -run '^TestBatchFinalizationPreservesFailureAcrossCleanup$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestBatchFinalizationPreservesFailureAcrossCleanup (0.00s)
    record007_regression_test.go:15: domain outcome lost primary failure: <nil>
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/batchcleanup	0.526s
FAIL
```
