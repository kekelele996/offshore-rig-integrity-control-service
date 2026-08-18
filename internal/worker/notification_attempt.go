package worker

type NotificationAttempt struct {
	Acknowledged bool
}

func (a *NotificationAttempt) Run(deliver func() error) error {
	if err := deliver(); err != nil {
		a.Acknowledged = false
		return err
	}
	a.Acknowledged = true
	return nil
}
