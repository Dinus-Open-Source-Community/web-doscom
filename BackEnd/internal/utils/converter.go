package utils

import (
	"reflect"
	"strings"
)

func StructToMap(structData any) map[string]any {
	data := make(map[string]any)
	value := reflect.ValueOf(structData)
	typeOfData := reflect.TypeOf(structData)

	// check if value is a pointer or Not
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		typeOfData = typeOfData.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		typeField := typeOfData.Field(i).Tag.Get("json")

		if field.IsZero() || typeField == "" || typeField == "-" {
			continue
		}

		typeField = strings.Split(typeField, ",")[0]
		data[typeField] = field.Interface()
	}

	return data
}
