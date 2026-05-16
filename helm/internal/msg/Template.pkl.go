// Code generated from Pkl module `helm.helm`. DO NOT EDIT.
package msg

import "github.com/apple/pkl-readers/helm/internal/msg/hookpolicy"

type Template interface {
	Request

	GetChart() string

	GetVersion() *string

	GetReleaseName() string

	GetNamespace() string

	GetKubeVersion() *string

	GetApiVersions() []string

	GetSkipSchemaValidation() bool

	GetHooks() hookpolicy.HookPolicy

	GetValuesJson() string
}

var _ Template = TemplateImpl{}

// Evaluate a Helm chart and "import" its output.
type TemplateImpl struct {
	Kind string `pkl:"kind"`

	// Identify the chart to template.
	//
	// May be in the format `<repo>/<chart>`, an OCI URI,
	// or an absolute path to a local chart directory or package.
	Chart string `pkl:"chart"`

	// If non-null, request a specific version of the chart.
	//
	// This is only relevant when [chart] is not a local path.
	Version *string `pkl:"version"`

	// Becomes the name of the Helm release.
	ReleaseName string `pkl:"releaseName"`

	// Equivalent to the `--namespace` flag of `helm template`.
	Namespace string `pkl:"namespace"`

	// Equivalent to the `--kube-version` flag of `helm template`.
	KubeVersion *string `pkl:"kubeVersion"`

	// Equivalent to the `--api-versions` flag of `helm template`.
	ApiVersions []string `pkl:"apiVersions"`

	// Equivalent to the `--skip-schema-validation` flag of `helm template`.
	SkipSchemaValidation bool `pkl:"skipSchemaValidation"`

	// Equivalent to the `--no-hooks` and `--skip-tests` flags of `helm template`.
	Hooks hookpolicy.HookPolicy `pkl:"hooks"`

	ValuesJson string `pkl:"valuesJson"`
}

func (rcv TemplateImpl) GetKind() string {
	return rcv.Kind
}

// Identify the chart to template.
//
// May be in the format `<repo>/<chart>`, an OCI URI,
// or an absolute path to a local chart directory or package.
func (rcv TemplateImpl) GetChart() string {
	return rcv.Chart
}

// If non-null, request a specific version of the chart.
//
// This is only relevant when [chart] is not a local path.
func (rcv TemplateImpl) GetVersion() *string {
	return rcv.Version
}

// Becomes the name of the Helm release.
func (rcv TemplateImpl) GetReleaseName() string {
	return rcv.ReleaseName
}

// Equivalent to the `--namespace` flag of `helm template`.
func (rcv TemplateImpl) GetNamespace() string {
	return rcv.Namespace
}

// Equivalent to the `--kube-version` flag of `helm template`.
func (rcv TemplateImpl) GetKubeVersion() *string {
	return rcv.KubeVersion
}

// Equivalent to the `--api-versions` flag of `helm template`.
func (rcv TemplateImpl) GetApiVersions() []string {
	return rcv.ApiVersions
}

// Equivalent to the `--skip-schema-validation` flag of `helm template`.
func (rcv TemplateImpl) GetSkipSchemaValidation() bool {
	return rcv.SkipSchemaValidation
}

// Equivalent to the `--no-hooks` and `--skip-tests` flags of `helm template`.
func (rcv TemplateImpl) GetHooks() hookpolicy.HookPolicy {
	return rcv.Hooks
}

func (rcv TemplateImpl) GetValuesJson() string {
	return rcv.ValuesJson
}
