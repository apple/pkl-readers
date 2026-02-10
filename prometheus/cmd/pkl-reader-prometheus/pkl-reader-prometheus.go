//===----------------------------------------------------------------------===//
// Copyright © 2026 Apple Inc. and the Pkl project authors. All rights reserved.
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
	_ "github.com/apple/pkl-readers/prometheus"
	"github.com/apple/pkl-readers/prometheus/internal"
	shared "github.com/apple/pkl-readers/shared/go"
)

var (
	Version   = "development"
	_, _, run = shared.New(shared.Spec{
		SchemeSuffix: "prometheus",
		Name:         "pkl-reader-prometheus",
		Short:        "Pkl External Reader for Prometheus",
		Long: `Pkl External Reader for Prometheus.

External Readers: https://pkl-lang.org/main/current/language-reference/index.html#external-readers

CLI configuration:
	--external-resource-reader reader+prometheus=pkl-reader-prometheus

PklProject configuration:
	evaluatorSettings {
		externalResourceReaders {
			["reader+prometheus"] {
				executable = "pkl-reader-prometheus"
			}
		}
	}
`,
		Version:           Version,
		VersionedPackages: []string{"github.com/prometheus/prometheus"},
	}, internal.Run)
)

func main() {
	run()
}
