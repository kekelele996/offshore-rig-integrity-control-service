package domain

import "errors"

func FinalizationOutcome(primary, cleanup error) error { return errors.Join(primary, cleanup) }
