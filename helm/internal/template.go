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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/apple/pkl-readers/helm/internal/msg"
	"helm.sh/helm/v4/pkg/action"
	"helm.sh/helm/v4/pkg/chart"
	"helm.sh/helm/v4/pkg/chart/v2/loader"
	"helm.sh/helm/v4/pkg/downloader"
	"helm.sh/helm/v4/pkg/getter"
	ri "helm.sh/helm/v4/pkg/release"
	release "helm.sh/helm/v4/pkg/release/v1"
)

func (r helmReader) template(req msg.Template) ([]byte, error) {
	var vals map[string]any
	if req.GetValuesJson() != "" {
		if err := json.Unmarshal([]byte(req.GetValuesJson()), &vals); err != nil {
			return nil, fmt.Errorf("failed to json decode values: %w", err)
		}
	}

	// adapted from https://github.com/helm/helm/blob/0ee89d2d4ee91d7edd21a9445f39f4eb0fed2973/pkg/cmd/template.go#L61
	cfg := new(action.Configuration)
	cfg.RegistryClient = r.registryClient
	client := action.NewInstall(cfg)
	client.DryRunStrategy = action.DryRunClient
	client.Replace = true
	client.IncludeCRDs = true
	client.ReleaseName = req.GetReleaseName()
	client.Namespace = req.GetNamespace()
	if version := req.GetVersion(); version != nil {
		client.Version = *version
	}

	slog.Debug(
		"run template",
		"release_name", client.ReleaseName,
		"namespace", client.Namespace,
		"chart", req.GetChart(),
		"version", req.GetVersion(),
	)

	providers := getter.All(r.settings)
	chartPath, err := client.LocateChart(req.GetChart(), r.settings)
	if err != nil {
		return nil, err
	}
	slog.Debug("located chart", "path", chartPath)

	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	chartAccessor, err := chart.NewAccessor(chartRequested)
	if err != nil {
		return nil, err
	}

	if err := checkIfInstallable(chartAccessor); err != nil {
		return nil, err
	}

	if chartAccessor.Deprecated() {
		slog.Warn("this chart is deprecated")
	}

	if reqs := chartAccessor.MetaDependencies(); reqs != nil {
		// If CheckDependencies returns an error, we have unfulfilled dependencies.
		// As of Helm 2.4.0, this is treated as a stopping condition:
		// https://github.com/helm/helm/issues/2209
		if err := action.CheckDependencies(chartRequested, reqs); err != nil {
			if client.DependencyUpdate {
				man := &downloader.Manager{
					Out:              os.Stderr,
					ChartPath:        chartPath,
					Keyring:          client.Keyring,
					SkipUpdate:       false,
					Getters:          providers,
					RepositoryConfig: r.settings.RepositoryConfig,
					RepositoryCache:  r.settings.RepositoryCache,
					Debug:            r.settings.Debug,
					RegistryClient:   client.GetRegistryClient(),
				}
				if err := man.Update(); err != nil {
					return nil, err
				}
				// Reload the chart with the updated Chart.lock file.
				if chartRequested, err = loader.Load(chartPath); err != nil {
					return nil, fmt.Errorf("failed reloading chart after repo update: %w", err)
				}
			} else {
				return nil, fmt.Errorf("an error occurred while checking for chart dependencies."+
					" You may need to run `helm dependency build` to fetch missing dependencies: %w", err)
			}
		}
	}

	rel1, err := client.RunWithContext(context.Background(), chartRequested, vals)
	if err != nil {
		return nil, fmt.Errorf("chart evaluation failed: %w", err)
	} else if rel1 == nil {
		return nil, nil
	}
	rel, err := releaserToV1Release(rel1)
	if err != nil {
		return nil, err
	}

	var manifests bytes.Buffer
	_, _ = fmt.Fprintln(&manifests, strings.TrimSpace(rel.Manifest))
	if !client.DisableHooks {
		for _, m := range rel.Hooks {
			if slices.Contains(m.Events, release.HookTest) {
				continue
			}
			_, _ = fmt.Fprintf(&manifests, "---\n# Source: %s\n%s\n", m.Path, m.Manifest)
		}
	}

	return manifests.Bytes(), nil
}

// adapted from https://github.com/helm/helm/blob/0ee89d2d4ee91d7edd21a9445f39f4eb0fed2973/pkg/cmd/install.go#L337
func checkIfInstallable(ch chart.Accessor) error {
	meta := ch.MetadataAsMap()

	switch meta["Type"] {
	case "", "application":
		return nil
	}
	return fmt.Errorf("%s charts are not installable", meta["Type"])
}

// adapted from https://github.com/helm/helm/blob/0ee89d2d4ee91d7edd21a9445f39f4eb0fed2973/pkg/cmd/root.go#L472
func releaserToV1Release(rel ri.Releaser) (*release.Release, error) {
	switch r := rel.(type) {
	case release.Release:
		return &r, nil
	case *release.Release:
		return r, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unsupported release type: %T", rel)
	}
}
