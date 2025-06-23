package errs

import "errors"

var (
	ErrUserExists     = errors.New("user already exists")
	ErrPassTooShort   = errors.New("password is too short")
	ErrPassTooLong    = errors.New("password is too long")
	ErrInvalidMail    = errors.New("invalid mail")
	ErrWrongLogOrPass = errors.New("wrong login or password")
)
