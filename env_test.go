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
	"os"
	"testing"
)

func TestGetVariableNameFromPrefixAndConfigPathEmptyPrefixSinglePathElement(t *testing.T) {
	result := getVariableNameFromPrefixAndConfigPath("", []string{"test"})
	expected := "TEST"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestGetVariableNameFromPrefixAndConfigPathEmptyPrefixLongPathElement(t *testing.T) {
	result := getVariableNameFromPrefixAndConfigPath("", []string{"Some", "basic", "TEST"})
	expected := "SOME_BASIC_TEST"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestGetVariableNameFromPrefixAndConfigPathNonEmptyPrefixLongPathElement(t *testing.T) {
	result := getVariableNameFromPrefixAndConfigPath("theprefix", []string{"Some", "basic", "TEST"})
	expected := "THEPREFIX_SOME_BASIC_TEST"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestGetExistingEnvironmentVariable(t *testing.T) {
	os.Setenv("SOME_EXISTING_VARIABLE", "some value")

	value, err := getVariableAsString("SOME_EXISTING_VARIABLE")
	if value != "some value" || err != nil {
		t.Errorf("could not get environment variable. got '%s'. error: '%s'", value, err.Error())
	}
}

func TestGetNonexistingEnvironmentVariable(t *testing.T) {
	value, err := getVariableAsString("SOME_NONEXISTING_VARIABLE")

	if value != "" || err == nil {
		t.Errorf("could access nonexisting variable. value: '%s'", value)
	}
}

func TestGetValidVariableAsInt(t *testing.T) {
	os.Setenv("SOME_VALID_INT_VALUE", "42")

	value, err := getVariableAsInt("SOME_VALID_INT_VALUE")

	if value != 42 || err != nil {
		t.Errorf("could not parse variable as int. got '%d'. error: %s", value, err.Error())
	}
}

func TestGetInvalidVariableAsInt(t *testing.T) {
	os.Setenv("SOME_INVALID_INT", "foo")

	value, err := getVariableAsInt("SOME_INVALID_INT")
	if err == nil {
		t.Errorf("could parse 'bla' as int. got: %d", value)
	}
}

func TestGetValidVariableAsBool(t *testing.T) {
	os.Setenv("SOME_VALID_BOOL", "y")

	value, err := getVariableAsBool("SOME_VALID_BOOL")
	if err != nil || !value {
		t.Errorf("could not get value as true. got '%t'. error: %s", value, err.Error())
	}
}

func TestGetInvalidVariableAsBool(t *testing.T) {
	os.Setenv("SOME_INVALID_BOOL", "bla")

	value, _ := getVariableAsBool("SOME_INVALID_BOOL")
	if value {
		t.Error("value was true, but shouldn't")
	}
}
