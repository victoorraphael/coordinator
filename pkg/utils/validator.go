package utils

import (
	"fmt"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"reflect"
)

func Validate(s interface{}) error {
	data, err := GetFields(s, "db")
	if err != nil {
		return err
	}

	for idx := range data {
		check := reflect.ValueOf(data[idx].Data).IsZero()
		if check {
			return errs.WrapError(errs.ErrFieldViolation, fmt.Sprintf("campo %s nao deve ser vazio", data[idx].Column))
		}
	}

	return nil
}

// RequiredFields check if value of fields are empty, based on tag "db"
func RequiredFields(s any, fields ...string) error {
	cols, err := GetFields(s, "db")
	if err != nil {
		return err
	}

	colsMap := cols.ToMap()
	for idx := range fields {
		if reflect.ValueOf(colsMap[fields[idx]]).IsZero() {
			return errs.WrapError(errs.ErrFieldViolation, fmt.Sprintf("campo %s nao deve ser vazio", fields[idx]))
		}
	}
	return nil
}
