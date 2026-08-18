# BUG_REPRO

## Bug 是什么

文件: internal/store/lease_store.go、internal/service/retry_service.go、internal/service/plan_service.go；符号: LeaseStore.Acquire、System.ReserveForRetry、System.RetryDisposition；机制: 多层包装用 %v 把 ErrLeaseConflict 的错误链转成纯文本，errors.Is 无法识别冲突，重试分类一路落到 fail。

## 如何触发

在项目根目录执行下面这条定向回归命令：

```bash
go test ./internal/service -run '^TestRetryRecognizesLeaseConflict$' -count=1
```

## 错误信息

埋错环境的真实输出如下：

```text
--- FAIL: TestRetryRecognizesLeaseConflict (0.00s)
    system_test.go:55: got fail
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/internal/service	0.271s
FAIL
```
