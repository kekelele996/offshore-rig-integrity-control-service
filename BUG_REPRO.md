# BUG_REPRO

## Bug 是什么

文件: internal/domain/notification_failure.go、internal/store/outbox_transaction.go、internal/service/notification_delivery.go、internal/worker/notification_attempt.go；符号: NotificationFailure.Unwrap、OutboxTransaction.Finish、DeliverNotification、NotificationAttempt.Run；机制: 通知提交失败在错误包装、事务 defer、服务命名返回和 worker 确认四个边界连续丢失，最终错误变 nil 且失败任务被确认。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test ./outboxdelivery -run '^TestNotificationDeliveryReturnsCommitFailure$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestNotificationDeliveryReturnsCommitFailure (0.00s)
    record009_regression_test.go:16: notification failure lost cause: commit: outbox commit failed
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/outboxdelivery	0.496s
FAIL
```
