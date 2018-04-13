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

import "reflect"

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
	v := reflect.ValueOf(s)
	field := v.FieldByName(fieldName)

	kind := field.Kind()

	switch kind {
	case reflect.Int:
		return field.Int()
	case reflect.Bool:
		return field.Bool()
	case reflect.Uint:
		return field.Uint()
	case reflect.Float32:
	case reflect.Float64:
		return field.Float()
	case reflect.String:
		return field.String()
	case reflect.Struct:
		return field.Interface()
	}

	return nil
}

func isFieldStruct(s interface{}, fieldName string) bool {
	v := reflect.ValueOf(s)
	field := v.FieldByName(fieldName)

	return field.Kind() == reflect.Struct
}
