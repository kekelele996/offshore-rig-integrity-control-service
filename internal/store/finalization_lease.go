package store

type FinalizationLease struct{ held bool }

func NewFinalizationLease() *FinalizationLease { return &FinalizationLease{held: true} }
func (l *FinalizationLease) Run(finalize func() error) (err error) {
	// The lease guards the asset for the duration of finalization. Whether
	// finalization succeeds or fails, the lease must be released so the next
	// shift can acquire the asset — a held lease on a failed batch is exactly
	// the stuck state we must avoid.
	defer func() {
		l.held = false
	}()
	return finalize()
}
func (l *FinalizationLease) Held() bool { return l.held }
