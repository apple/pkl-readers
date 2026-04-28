// Code generated from Pkl module `prometheus.promql`. DO NOT EDIT.
package promql

import "github.com/apple/pkl-readers/prometheus/internal/msg"

type Parse interface {
	msg.Request

	GetQuery() string

	GetEnableFeatures() []string
}

var _ Parse = ParseImpl{}

// Parse a [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) expression.
type ParseImpl struct {
	Kind string `pkl:"kind"`

	// PromQL query to parse.
	Query string `pkl:"query"`

	// Prometheus [feature flags](https://prometheus.io/docs/prometheus/latest/feature_flags/) to enable.
	//
	// [Feature] documents supported features.
	// Unrecognized features result in a warning and are ignored.
	EnableFeatures []string `pkl:"enableFeatures"`
}

func (rcv ParseImpl) GetKind() string {
	return rcv.Kind
}

// PromQL query to parse.
func (rcv ParseImpl) GetQuery() string {
	return rcv.Query
}

// Prometheus [feature flags](https://prometheus.io/docs/prometheus/latest/feature_flags/) to enable.
//
// [Feature] documents supported features.
// Unrecognized features result in a warning and are ignored.
func (rcv ParseImpl) GetEnableFeatures() []string {
	return rcv.EnableFeatures
}
