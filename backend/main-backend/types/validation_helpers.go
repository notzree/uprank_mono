package types

import "reflect"

func findNilFields(v interface{}) []string {
	val := reflect.ValueOf(v).Elem()
	nilArray := []string{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Check if the field is a pointer, interface, map, slice, or channel
		switch field.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan:
			if field.IsNil() {
				nilArray = append(nilArray, fieldType.Name)
			}
		}
	}

	return nilArray
}

func getNumFields(v interface{}) int {
	return reflect.ValueOf(v).Elem().NumField()
}
