// Copyright 2018 Jannis Fink
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enconf

import (
	"errors"
	"fmt"
	"reflect"
)

func getShallowFieldNamesInStruct(s interface{}) []string {
	result := []string{}

	r := reflect.ValueOf(s).Elem()
	typeOfT := r.Type()
	for i := 0; i < r.NumField(); i++ {
		result = append(result, typeOfT.Field(i).Name)
	}

	return result
}

func getValueOfStructField(s interface{}, fieldName string) interface{} {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)

	kind := field.Kind()

	switch kind {
	case reflect.Int:
		return field.Int()
	case reflect.Bool:
		return field.Bool()
	case reflect.String:
		return field.String()
	case reflect.Struct:
		return field.Interface()
	}

	return nil
}

func setValueOfStructField(s interface{}, fieldName string, fieldValue interface{}) error {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)

	if !field.CanAddr() || !field.CanSet() {
		return errors.New(fmt.Sprintf("cannot set field '%s'", fieldName))
	}

	switch getTypeOfField(s, fieldName) {
	case reflect.Int:
		field.SetInt(fieldValue.(int64))
		return nil
	case reflect.Bool:
		field.SetBool(fieldValue.(bool))
		return nil
	case reflect.String:
		field.SetString(fieldValue.(string))
		return nil
	default:
		return errors.New(fmt.Sprintf("field '%s' has none of the allowed formats (int, bool, string)", fieldName))
	}
}

func isFieldStruct(s interface{}, fieldName string) bool {
	return getTypeOfField(s, fieldName) == reflect.Struct
}

func getTypeOfField(s interface{}, fieldName string) reflect.Kind {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)
	return field.Kind()
}
