// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2020 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package downloads

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed tarball_list.json
var _tarballList []byte

var DefaultTarballRegistry TarballCollection

func init() {

	err := json.Unmarshal(_tarballList, &DefaultTarballRegistry)
	if err != nil {
		fmt.Printf("[FATAL] error decoding JSON file into default tarball registry: %s\n", err)
		os.Exit(1)
	}
}
