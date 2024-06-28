package types

import "reflect"

func findNilFields(v interface{}) []string {
	val := reflect.ValueOf(v).Elem()
	nilArray := []string{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.IsNil() {
			nilArray = append(nilArray, val.Type().Field(i).Name)
		}
	}

	return nilArray
}

func getNumFields(v interface{}) int {
	return reflect.ValueOf(v).Elem().NumField()
}
