package worker

type NotificationAttempt struct{ Acknowledged bool }

func (a *NotificationAttempt) Run(deliver func() error) (err error) {
	defer func() { a.Acknowledged = true }()
	return deliver()
}
