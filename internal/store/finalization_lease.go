package store

type FinalizationLease struct{ held bool }

func NewFinalizationLease() *FinalizationLease { return &FinalizationLease{held: true} }
func (l *FinalizationLease) Run(finalize func() error) (err error) {
	defer func() { l.held = false }()
	return finalize()
}
func (l *FinalizationLease) Held() bool { return l.held }
