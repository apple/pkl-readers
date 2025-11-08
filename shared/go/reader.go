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
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/apple/pkl-go/pkl"
)

func Run(ctx context.Context, spec Spec, opts ...func(options *pkl.ExternalReaderClientOptions)) error {
	runtime, err := pkl.NewExternalReaderClient(opts...)
	if err != nil {
		return err
	}
	slog.InfoContext(ctx, "starting external reader", "name", spec.Name, "version", spec.Version)
	if err := runtime.Run(); err != nil {
		return err
	}
	slog.InfoContext(ctx, "stopping external reader", "name", spec.Name, "version", spec.Version)
	return nil
}

func (s Spec) CheckPackageVersion(uri url.URL) error {
	if !s.enablePackageVersionCheck {
		return nil
	}

	packageVersion := uri.Query().Get("packageVersion")
	switch packageVersion {
	case "":
		return errors.New("read uri did not include expected packageVersion query parameter")
	case s.Version:
		return nil
	default:
		return fmt.Errorf("package and reader version mismatch: package has version '%s' and reader has version '%s'", packageVersion, s.Version)
	}
}
