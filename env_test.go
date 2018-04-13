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

import "testing"

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
