# BUG_REPRO

## Bug 是什么

文件: internal/domain/manifest_snapshot.go、internal/store/manifest_snapshot_store.go、internal/service/manifest_snapshot_service.go、internal/worker/manifest_snapshot_dispatch.go；符号: CopyManifestSnapshot、ManifestSnapshotStore.Save/Load、EmergencyManifestView、DispatchManifest；机制: 清单在领域复制、仓储读写、服务扩展和异步派发四个所有权边界都保留了可变切片与嵌套 map 的共享引用，后续调用会改写既有快照。

## 如何触发

在项目根目录执行下面这条回归命令：

```bash
go test ./manifestownership -race -run '^TestManifestSnapshotHasIndependentMemory$' -count=1
```

## 错误信息

埋错环境的真实输出如下（退出码 1）：

```text
--- FAIL: TestManifestSnapshotHasIndependentMemory (0.00s)
    record001_regression_test.go:29: domain snapshot changed with caller input: {PlanID:rig-17 Zones:[mutated-input south] Contacts:[control-room] RequiredChecks:map[north:[mutated-check corrosion]]}
FAIL
FAIL	github.com/kekelele996/offshore-rig-integrity-control-service/manifestownership	0.523s
FAIL
```
