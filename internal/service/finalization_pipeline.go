package service

import "errors"

func FinalizeInspection(finalize, cleanup func() error) (err error) {
	defer func() { err = errors.Join(err, cleanup()) }()
	err = finalize()
	return
}
