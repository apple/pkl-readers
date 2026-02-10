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
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/apple/pkl-go/pkl"
	"github.com/apple/pkl-readers/prometheus/internal/msg"
	"github.com/apple/pkl-readers/prometheus/internal/msg/promql"
	shared "github.com/apple/pkl-readers/shared/go"
)

type Options struct{}

func Run(ctx context.Context, spec shared.Spec, _ *Options) error {
	reader := prometheusReader{
		Spec: spec,
	}

	return shared.Run(ctx, spec,
		pkl.WithExternalClientResourceReader(reader),
	)
}

type prometheusReader struct {
	shared.Spec
}

func (r prometheusReader) Read(uri url.URL) ([]byte, error) {
	var req msg.Request
	if err := r.DecodeRequest(uri, &req); err != nil {
		return nil, err
	}

	slog.Debug("received request", "kind", req.GetKind())

	switch reqType := req.(type) {
	case promql.Parse:
		return r.parse(reqType)
	default:
		return nil, fmt.Errorf("unrecognized action '%s'", uri.Host)
	}
}
