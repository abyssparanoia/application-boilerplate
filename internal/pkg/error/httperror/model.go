package httperror

import (
	"fmt"

	"github.com/pkg/errors"
)

// InternalError ... base error
type InternalError string

func (e InternalError) Error() string {
	return string(e)
}

// New ... new error
func (e InternalError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e InternalError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e InternalError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e InternalError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// ServiceUnavailableError ... service unavailable error
type ServiceUnavailableError string

func (e ServiceUnavailableError) Error() string {
	return string(e)
}

// New ... new error
func (e ServiceUnavailableError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e ServiceUnavailableError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e ServiceUnavailableError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e ServiceUnavailableError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// As ... as method
func (e ServiceUnavailableError) As(target interface{}) bool {
	if _, ok := target.(**ServiceUnavailableError); ok {
		return true
	}
	return false
}

// BadRequestError ... bad request error
type BadRequestError string

func (e BadRequestError) Error() string {
	return string(e)
}

// New ... new error
func (e BadRequestError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e BadRequestError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e BadRequestError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e BadRequestError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// As ... as method
func (e BadRequestError) As(target interface{}) bool {
	if _, ok := target.(**BadRequestError); ok {
		return true
	}
	return false
}

// ForbiddenError ... forbidden error
type ForbiddenError string

func (e ForbiddenError) Error() string {
	return string(e)
}

// New ... new error
func (e ForbiddenError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e ForbiddenError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e ForbiddenError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e ForbiddenError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// As ... as method
func (e ForbiddenError) As(target interface{}) bool {
	if _, ok := target.(**ForbiddenError); ok {
		return true
	}
	return false
}

// NotFoundError ... not found error
type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}

// New ... new error
func (e NotFoundError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e NotFoundError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e NotFoundError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e NotFoundError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// As ... as method
func (e NotFoundError) As(target interface{}) bool {
	if _, ok := target.(**NotFoundError); ok {
		return true
	}
	return false
}

// UnauthorizedError ... not found error
type UnauthorizedError string

func (e UnauthorizedError) Error() string {
	return string(e)
}

// New ... new error
func (e UnauthorizedError) New() error {
	return errors.Wrap(e, "")
}

// Errorf ... errorf
func (e UnauthorizedError) Errorf(format string, args ...interface{}) error {
	return errors.Wrapf(e, format, args...)
}

// Wrap ... wrap error
func (e UnauthorizedError) Wrap(err error) error {
	if err == nil {
		return e.New()
	}
	return errors.Wrap(e, err.Error())
}

// Wrapf ... wrapf
func (e UnauthorizedError) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return e.Errorf(format, args...)
	}
	msg := fmt.Sprintf(format, args...)
	return errors.Wrapf(e, "err: %s; %s", err, msg)
}

// As ... as method
func (e UnauthorizedError) As(target interface{}) bool {
	if _, ok := target.(**UnauthorizedError); ok {
		return true
	}
	return false
}
