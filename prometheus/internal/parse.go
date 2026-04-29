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
	"log/slog"

	"github.com/apple/pkl-readers/prometheus/internal/msg/promql"
	"github.com/prometheus/prometheus/promql/parser"
)

func toParserOptions(features []string) parser.Options {
	var opt parser.Options
	for _, o := range features {
		switch o {
		case "promql-experimental-functions":
			opt.EnableExperimentalFunctions = true
		case "promql-duration-expr":
			opt.ExperimentalDurationExpr = true
		case "promql-extended-range-selectors":
			opt.EnableExtendedRangeSelectors = true
		case "promql-binop-fill-modifiers":
			opt.EnableBinopFillModifiers = true
		default:
			slog.Warn("unknown feature in enableFeatures", "feature_name", o)
		}
	}
	return opt
}

func (r prometheusReader) parse(req promql.Parse) ([]byte, error) {
	p := parser.NewParser(toParserOptions(req.GetEnableFeatures()))
	if _, err := p.ParseExpr(req.GetQuery()); err != nil {
		// intentionally return error as resource content and not an error
		// let calling pkl code determine how to handle this
		return []byte(err.Error()), nil
	}

	// on success, return empty resource content
	return nil, nil
}
