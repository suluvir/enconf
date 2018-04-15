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
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var validTrueValues = []string{
	"true",
	"1",
	"y",
}

func getVariableNameFromPrefixAndConfigPath(prefix string, configPath []string) string {
	result := ""

	for _, p := range append([]string{prefix}, configPath...) {
		result += fmt.Sprintf("%s_", strings.ToUpper(p))
	}

	return strings.Trim(result, "_")
}

func getVariableAsString(name string) (string, error) {
	if value, ok := os.LookupEnv(name); ok {
		return value, nil
	}
	return "", fmt.Errorf("cannot find environment variable '%s'", name)
}

func getVariableAsInt(name string) (int, error) {
	value, err := getVariableAsString(name)
	if err != nil {
		return 0, err
	}

	result, parseErr := strconv.Atoi(value)
	if parseErr != nil {
		return 0, fmt.Errorf("cannot parse variable '%s' as int: %s", name, parseErr.Error())
	}
	return result, nil
}

func getVariableAsBool(name string) (bool, error) {
	value, err := getVariableAsString(name)
	if err != nil {
		return false, err
	}
	value = strings.ToLower(value)
	for _, trueValue := range validTrueValues {
		if trueValue == value {
			return true, nil
		}
	}
	return false, nil
}

func getVariableAsKind(name string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Bool:
		return getVariableAsBool(name)
	case reflect.String:
		return getVariableAsString(name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return getVariableAsInt(name)
	}
	return nil, nil
}
