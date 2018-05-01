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

type shallowStruct struct {
	SomeInt    int
	SomeBool   bool
	SomeString string
}

type deepStruct struct {
	SomeInt    int
	SomeStruct struct{
		SomeNested bool
	}
}

func TestLoadConfigurationWithPrefixShallow(t *testing.T) {
	a := shallowStruct{}

	os.Setenv("SULUVIR_SOMEINT", "15")
	os.Setenv("SULUVIR_SOMEBOOL", "true")
	os.Setenv("SULUVIR_SOMESTRING", "moin")

	err := LoadConfigurationWithPrefix("SULUVIR", &a)

	if err != nil {
		t.Errorf("loading configuration failed. Error: %s", err.Error())
	}

	if !a.SomeBool || a.SomeInt != 15 || a.SomeString != "moin" {
		t.Error("loading configuration failed.")
	}
}

func TestLoadConfigurationWithPrefixDeep(t *testing.T) {
	a := deepStruct{}

	os.Setenv("SULUVIR_SOMEINT", "42")
	os.Setenv("SULUVIR_SOMESTRUCT_SOMENESTED", "1")

	err := LoadConfigurationWithPrefix("SULUVIR", &a)

	if err != nil {
		t.Errorf("loading configuration failed. Error: %s", err.Error())
	}

	if a.SomeInt != 42 || !a.SomeStruct.SomeNested {
		t.Error("loading configuration failed.")
	}
}
