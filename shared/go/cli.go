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

package shared

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var enableDebugLogging bool

type Spec struct {
	Name                      string
	Short                     string
	Long                      string
	Version                   string
	VersionedPackages         []string
	enablePackageVersionCheck bool
}

type version struct {
	name    string
	version string
}

func New[T any](spec Spec, run func(ctx context.Context, spec Spec, opts *T) error) (*pflag.FlagSet, *T, func()) {
	var opts T
	command := &cobra.Command{
		Use:   spec.Name,
		Short: spec.Short,
		Long:  spec.Long,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			if enableDebugLogging {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
			if err := run(cmd.Context(), spec, &opts); err != nil {
				slog.ErrorContext(cmd.Context(), "execution failed", "error", err)
			}
		},
	}
	command.Flags().BoolVar(&spec.enablePackageVersionCheck, "package-version-check", true, "Disable checking the package version is the same as the reader version")
	command.Flags().BoolVar(&enableDebugLogging, "debug", false, "Enable debug level logging")
	_ = command.Flags().MarkHidden("package-version-check")

	versionCommand := &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Print the version of %s and major dependencies", spec.Name),
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			printVersion(spec)
		},
	}
	command.AddCommand(versionCommand)

	return command.Flags(), &opts, func() {
		if err := command.Execute(); err != nil {
			slog.Error("execution failed", "error", err)
		}
	}
}

func printVersion(spec Spec) {
	var versions []version
	longestVersion := 0
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("failed to read build info")
	}

	if info.Main.Version != "" && info.Main.Version != "(devel)" && spec.Version == "development" {
		spec.Version = strings.TrimPrefix(info.Main.Version, "v")
	}
	longestVersion = len(spec.Version)
	versions = append(versions, version{spec.Name, spec.Version})
	versions = append(versions, version{"go", strings.TrimPrefix(info.GoVersion, "go")})
Pkg:
	for _, pkg := range spec.VersionedPackages {
		longestVersion = max(longestVersion, len(pkg))
		for _, dep := range info.Deps {
			if pkg == dep.Path {
				versions = append(versions, version{pkg, strings.TrimPrefix(dep.Version, "v")})
				continue Pkg
			}
		}
		panic("could not determine version for dep " + pkg)
	}

	for _, ver := range versions {
		fmt.Printf("%-*s %s\n", longestVersion, ver.name, ver.version)
	}
}
