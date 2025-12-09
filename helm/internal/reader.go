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

package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/apple/pkl-go/pkl"
	"github.com/apple/pkl-readers/helm/internal/msg"
	shared "github.com/apple/pkl-readers/shared/go"
	"helm.sh/helm/v4/pkg/cli"
	"helm.sh/helm/v4/pkg/registry"
)

type Options struct{}

func Run(ctx context.Context, spec shared.Spec, _ *Options) error {
	registryClient, err := registry.NewClient(
		registry.ClientOptDebug(false),
		registry.ClientOptEnableCache(true),
		registry.ClientOptWriter(os.Stderr),
	)
	if err != nil {
		return err
	}

	reader := helmReader{
		Spec:           spec,
		registryClient: registryClient,
		settings:       cli.New(),
	}

	return shared.Run(ctx, spec,
		pkl.WithExternalClientResourceReader(reader),
	)
}

type helmReader struct {
	shared.Spec
	registryClient *registry.Client
	settings       *cli.EnvSettings
}

func (r helmReader) Read(uri url.URL) ([]byte, error) {
	var req msg.Request
	if err := r.DecodeRequest(uri, &req); err != nil {
		return nil, err
	}

	slog.Debug("received request", "kind", req.GetKind())

	switch reqType := req.(type) {
	case msg.Template:
		return r.template(reqType)
	default:
		return nil, fmt.Errorf("unrecognized action '%s'", uri.Host)
	}
}
