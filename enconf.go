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
	return nil
}
