package validation

import (
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/glueerr"
	"github.com/go-playground/validator"
)

var v = validator.New()

func Struct(s interface{}) error {
	if err := v.Struct(s); err != nil {
		return glueerr.BadRequestErr.Wrap(err)
	}
	return nil
}
