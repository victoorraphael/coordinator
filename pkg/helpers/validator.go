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

	for idx := range data {
		check := reflect.ValueOf(data[idx].Data).IsZero()
		if check {
			return errs.WrapError(errs.ErrFieldViolation, fmt.Sprintf("campo %s nao deve ser vazio", data[idx].Column))
		}
	}

	return nil
}
