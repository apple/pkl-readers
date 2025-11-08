Copyright © 2025 Apple Inc. and the Pkl project authors

The pkl-readers project's binaries include libraries that may be distributed under a different license.

{{ range $index, $value := . }}
---
{{ .Name }}
{{ .LicenseName }} - {{ .LicenseURL }}

***

{{ .LicenseText }}
{{ end }}
