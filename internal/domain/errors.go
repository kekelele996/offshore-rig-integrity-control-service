package domain

import "errors"

var (
    ErrLeaseConflict = errors.New("asset lease conflict")
    ErrPolicyUnavailable = errors.New("policy provider unavailable")
    ErrUnsafeRelease = errors.New("unsafe release")
    ErrCancelled = errors.New("inspection was cancelled")
    ErrFinalization = errors.New("batch finalization failed")
    ErrInvalidState = errors.New("invalid plan state")
)
