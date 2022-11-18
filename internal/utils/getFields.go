package utils

import (
	"database/sql/driver"
	"reflect"
)

// ColumnDataPair describes a piece of data that is stored in a database table column.
type ColumnDataPair struct {
	Column string
	Data   interface{}
}

// GetFields returns an array of ColumnDataPair which describes a database row.
// It uses the db struct tab to get the column name.
func GetFields(s interface{}) ([]ColumnDataPair, error) {
	var row []ColumnDataPair
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := field.Tag.Get("db")
		if col == "" {
			col = field.Name
		}

		val, err := driver.DefaultParameterConverter.ConvertValue(v.Field(i).Interface())
		if err != nil {
			return row, err
		}

		row = append(row, ColumnDataPair{
			Column: col,
			Data:   val,
		})
	}

	return row, nil
}
