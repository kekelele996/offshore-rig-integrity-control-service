package domain

type NotificationFailure struct {
	Stage string
	Cause error
}

func (e *NotificationFailure) Error() string { return e.Stage + ": " + e.Cause.Error() }
func (e *NotificationFailure) Unwrap() error { return nil }
