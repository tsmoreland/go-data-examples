package shared

import "errors"

var (
	ErrNotImplemented  = errors.New("not implemented")
	ErrDuplicate       = errors.New("item already exists")
	ErrNotFound        = errors.New("no match found")
	ErrInvalidArgument = errors.New("one or more arguments is invalid")
	ErrDeleteFailed    = errors.New("delete failed, no matching rows found")
	ErrUnableToUpdate  = errors.New("unable to update, record does not exist")
)
