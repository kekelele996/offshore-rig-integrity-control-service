package domain

import "fmt"

type NotificationFailure struct {
	Stage string
	Cause error
}

func (e *NotificationFailure) Error() string {
	return fmt.Sprintf("notification %s: %v", e.Stage, e.Cause)
}

func (e *NotificationFailure) Unwrap() error {
	return e.Cause
}
