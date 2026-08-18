package worker

import "github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain"

func DispatchManifest(snapshot domain.ManifestSnapshot) <-chan domain.ManifestSnapshot {
	out := make(chan domain.ManifestSnapshot, 1)
	go func() {
		out <- domain.CopyManifestSnapshot(snapshot)
		close(out)
	}()
	return out
}
