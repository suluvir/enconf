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

// LoadConfiguration loads the configuration from the environment and saves the configuration in the given
// structure. See `LoadConfigurationWithPrefix` for further info regarding the names of the environment
// variables parsed.
func LoadConfiguration(configStruct interface{}) error {
	return LoadConfigurationWithPrefix("", configStruct)
}

// LoadConfigurationWithPrefix does the same as `LoadConfiguration`, but allows for a use of prefixes. The given
// prefix string will me transformed into uppercase as the struct field names for the configuration struct.
func LoadConfigurationWithPrefix(prefix string, configStruct interface{}) error {
	return loadConfigurationRecursively(prefix, configStruct, []string{})
}

func loadConfigurationRecursively(prefix string, configStruct interface{}, configPath []string) error {
	fields := getShallowFieldNamesInStruct(configStruct)
	for _, field := range fields {
		newConfigPath := append(configPath, field)

		if isFieldStruct(configStruct, field) {
			newConfigStruct := getValueOfStructField(configStruct, field)
			err := loadConfigurationRecursively(prefix, newConfigStruct, newConfigPath)
			if err != nil {
				return err
			}
		} else {
			err := loadConfigurationVariable(prefix, configStruct, field, newConfigPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func loadConfigurationVariable(prefix string, configStruct interface{}, field string, configPath []string) error {
	varName := getVariableNameFromPrefixAndConfigPath(prefix, configPath)
	value, getErr := getVariableAsKind(varName, getKindOfField(configStruct, field))
	if getErr != nil {
		return getErr
	}
	setErr := setValueOfStructField(configStruct, field, value)
	if setErr != nil {
		return setErr
	}

	return nil
}
