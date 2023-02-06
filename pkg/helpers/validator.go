package helpers

import (
	"fmt"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"github.com/victoorraphael/coordinator/pkg/utils"
	"reflect"
)

func Validate(s interface{}) error {
	data, err := utils.GetFields(s, "db")
	if err != nil {
		return err
	}

	var errField error
	for idx := range data {
		check := reflect.ValueOf(data[idx].Data).IsZero()
		if check {
			errField = errs.WrapError(errs.ErrFieldViolation, fmt.Sprintf("%s should not be empty", data[idx].Column))
		}
	}

	return errField
}
