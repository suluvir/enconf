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
	"reflect"
	"testing"
)

type testStructA struct {
	A              string
	SomeOtherField int
}

type testStructB struct {
	B     string
	Other struct {
		B2 int
	}
}

func TestGetShallowFieldNamesInStruct(t *testing.T) {
	v := testStructA{}
	result := getShallowFieldNamesInStruct(&v)
	expected := []string{"A", "SomeOtherField"}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected slices to be equal; expected '%s', got '%s'", expected, result)
	}
}

func TestGetValueOfStructField(t *testing.T) {
	v := testStructA{
		A:              "test",
		SomeOtherField: 42,
	}

	resultGetA := getValueOfStructField(v, "A")
	expectedGetA := "test"
	if resultGetA != expectedGetA {
		t.Errorf("Expected to extract '%s', got '%s'", expectedGetA, resultGetA)
	}

	resultGetOther := getValueOfStructField(v, "SomeOtherField")
	expectedGetOther := int64(42)
	if resultGetOther != expectedGetOther {
		t.Errorf("Expected to extract '%d', got '%d'", expectedGetOther, resultGetOther)
	}
}

func TestIsFieldStruct(t *testing.T) {
	v := testStructB{
		B:     "test",
		Other: struct{ B2 int }{B2: 1337},
	}

	if isFieldStruct(v, "B") {
		t.Error("Expected field 'B' not to be struct, but was")
	}

	if !isFieldStruct(v, "Other") {
		t.Error("Expected field 'Other' to be struct, but wasn't")
	}
}

//func TestSetFieldValue(t *testing.T) {
//	v := testStructB{}
//
//	err := setValueOfStructField(v, "B", "test")
//	if err != nil {
//		t.Errorf("could not set field value. error: %s", err.Error())
//	}
//}
