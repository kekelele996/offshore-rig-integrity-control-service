# BUG_REPRO

## Bug 是什么

文件: internal/domain/handover_roster.go、internal/store/handover_history.go、internal/service/handover_filter.go、internal/worker/handover_assignment.go；符号: ActiveHandoverRoster、HandoverHistory.Visible、ContingencyRoster、AssignBackupCrew；机制: 四层过滤逻辑都用 roster[:0] 原地压缩共享切片，过滤结果覆盖调用方和历史记录持有的原始交接名单。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test ./handoverroster -run '^TestHandoverContingencySnapshotRemainsDetached$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestHandoverContingencySnapshotRemainsDetached (0.00s)
    record006_regression_test.go:17: domain compaction changed retained roster: [marine subsea medic ]
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/handoverroster	0.515s
FAIL
```
