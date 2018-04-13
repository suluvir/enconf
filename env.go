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
	"strings"
)

func getVariableNameFromPrefixAndConfigPath(prefix string, configPath []string) string {
	result := ""

	for _, p := range append([]string{prefix}, configPath...) {
		result += fmt.Sprintf("%s_", strings.ToUpper(p))
	}

	return strings.Trim(result, "_")
}
