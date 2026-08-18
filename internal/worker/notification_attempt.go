package worker

type NotificationAttempt struct{ Acknowledged bool }

func (a *NotificationAttempt) Run(deliver func() error) (err error) {
	err = deliver()
	if err != nil {
		return err
	}
	a.Acknowledged = true
	return nil
}
