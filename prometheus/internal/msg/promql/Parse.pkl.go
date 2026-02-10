// Code generated from Pkl module `pkl.prometheus.promql`. DO NOT EDIT.
package promql

import "github.com/apple/pkl-readers/prometheus/internal/msg"

type Parse interface {
	msg.Request

	GetQuery() string
}

var _ Parse = ParseImpl{}

// Parse a [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) expression.
type ParseImpl struct {
	Kind string `pkl:"kind"`

	// PromQL query to parse.
	Query string `pkl:"query"`
}

func (rcv ParseImpl) GetKind() string {
	return rcv.Kind
}

// PromQL query to parse.
func (rcv ParseImpl) GetQuery() string {
	return rcv.Query
}
