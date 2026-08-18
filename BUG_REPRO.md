# BUG_REPRO

## Bug 是什么

文件: internal/domain/finding_window.go、internal/store/risk_cache.go、internal/service/risk_cache_service.go、internal/worker/risk_cache_fanout.go；符号: FindingWindow.Snapshot、RiskCache.Read、RiskProjector.Project、FanoutRisk；机制: 领域窗口、缓存读取、服务投影和 worker fanout 都把可复用 scratch buffer 直接交给消费者，后续复用或另一个消费者写入会覆盖先前视图。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test ./riskbuffers -run '^TestRiskCacheReadersReceiveIndependentSnapshots$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestRiskCacheReadersReceiveIndependentSnapshots (0.00s)
    record010_regression_test.go:16: domain window reused retained view: [{Code:weather Severity:high Message:}]
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/riskbuffers	0.802s
FAIL
```
