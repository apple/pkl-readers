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
	"encoding/base64"
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

// DecodeRequest parses the `request` query parameter out of the provided URI, base64 decodes it, then unmarshals the pkl-binary data into `v`.
// Also, when s.enablePackageVersionCheck is true it checks the `packageVersion` query parameter against s.Version.
func (s Spec) DecodeRequest(uri url.URL, v any) error {
	req := uri.Query().Get("request")
	if req == "" {
		return errors.New("no request parameter found in uri")
	}

	reqBuf, err := base64.RawURLEncoding.DecodeString(req)
	if err != nil {
		return fmt.Errorf("unable to decode request: %w", err)
	}

	if s.enablePackageVersionCheck {
		packageVersion := uri.Query().Get("packageVersion")
		switch packageVersion {
		case s.Version: // success case
			break
		case "":
			return errors.New("read uri did not include expected packageVersion query parameter")
		default:
			return fmt.Errorf("package and reader version mismatch: package has version '%s' and reader has version '%s'", packageVersion, s.Version)
		}
	}

	if err := pkl.Unmarshal(reqBuf, v); err != nil {
		return fmt.Errorf("unable to unmarshal request: %w", err)
	}
	return nil
}

var _ pkl.Reader = (*Spec)(nil)

func (s Spec) Scheme() string                                    { return s.scheme }
func (s Spec) IsGlobbable() bool                                 { return false }
func (s Spec) HasHierarchicalUris() bool                         { return true }
func (s Spec) ListElements(_ url.URL) ([]pkl.PathElement, error) { return nil, nil }
