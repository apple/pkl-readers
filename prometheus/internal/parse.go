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

package internal

import (
	"github.com/apple/pkl-readers/prometheus/internal/msg/promql"
	"github.com/prometheus/prometheus/promql/parser"
)

func (r prometheusReader) parse(req promql.Parse) ([]byte, error) {
	p := parser.NewParser(req.GetQuery())
	if _, err := p.ParseExpr(); err != nil {
		// intentionally return error as resource content and not an error
		// let calling pkl code determine how to handle this
		return []byte(err.Error()), nil
	}

	// on success, return empty resource content
	return nil, nil
}
