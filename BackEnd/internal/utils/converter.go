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
		fieldType := typeOfData.Field(i)

		// Coba ambil tag json, kalau gak ada coba tag form
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = fieldType.Tag.Get("form")
		}

		// Jika masih kosong tapi field ini punya data, gunakan nama field asli (lowercase) sebagai fallback
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}

		// Skip jika tag di-set "-" atau jika field kosong (Zero Value)
		// Ini penting untuk partial update agar string kosong tidak menimpa data di database
		if tag == "-" || field.IsZero() {
			continue
		}

		tag = strings.Split(tag, ",")[0]

		// SKIP field jika tipenya adalah pointer ke multipart.FileHeader (untuk gambar)
		// karena ini bukan kolom database
		if fieldType.Type.String() == "*multipart.FileHeader" {
			continue
		}

		data[tag] = field.Interface()
	}

	return data
}
