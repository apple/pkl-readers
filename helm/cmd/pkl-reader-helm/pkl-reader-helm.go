//===----------------------------------------------------------------------===//
// Copyright © 2025 Apple Inc. and the Pkl project authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//===----------------------------------------------------------------------===//

package main

import (
	_ "github.com/apple/pkl-readers/helm"
	"github.com/apple/pkl-readers/helm/internal"
	shared "github.com/apple/pkl-readers/shared/go"
)

var (
	Version   = "development"
	_, _, run = shared.New(shared.Spec{
		Name:  "pkl-reader-helm",
		Short: "Pkl External Reader for Helm charts",
		Long: `Pkl External Reader for Helm charts.

External Readers: https://pkl-lang.org/main/current/language-reference/index.html#external-readers

CLI configuration:
	--external-resource-reader reader+helm=pkl-reader-helm

PklProject configuration:
	evaluatorSettings {
		externalResourceReaders {
			["reader+helm"] {
				executable = "pkl-reader-helm"
			}
		}
	}
`,
		Version:           Version,
		VersionedPackages: []string{"helm.sh/helm/v4"},
	}, internal.Run)
)

func main() {
	run()
}
