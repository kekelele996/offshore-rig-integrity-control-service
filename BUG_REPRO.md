# BUG_REPRO

## Bug 是什么

文件: internal/domain/provider_presence.go、internal/store/policy_registry.go、internal/service/policy_adapter.go、internal/worker/policy_dispatch.go；符号: ProviderHandle.Available、PolicyRegistry.Register/Has、EvaluateProvider、AcceptPolicyJob；机制: typed-nil 指针装入接口后接口值非 nil，四个跨层可用性边界都仅做接口判空，使未初始化策略被注册、评估并派发。

## 如何触发

在项目根目录执行下面这条定向回归命令：

```bash
go test ./internal/service -run '^TestTypedNilPolicyRejectedAtEveryBoundary$' -count=1
```

## 错误信息

埋错环境的真实输出如下：

```text
--- FAIL: TestTypedNilPolicyRejectedAtEveryBoundary (0.00s)
    record004_regression_test.go:24: domain handle reported a typed-nil provider as available
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/internal/service	0.314s
FAIL
```
