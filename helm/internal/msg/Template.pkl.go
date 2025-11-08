// Code generated from Pkl module `pkl.helm.helm`. DO NOT EDIT.
package msg

type Template interface {
	Request

	GetChart() string

	GetVersion() *string

	GetReleaseName() string

	GetNamespace() string

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

func (rcv TemplateImpl) GetValuesJson() string {
	return rcv.ValuesJson
}
