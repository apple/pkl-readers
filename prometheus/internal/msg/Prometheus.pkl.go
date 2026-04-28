// Code generated from Pkl module `prometheus.prometheus`. DO NOT EDIT.
package msg

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Prometheus struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Prometheus
func LoadFromPath(ctx context.Context, path string) (ret Prometheus, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return ret, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Prometheus
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Prometheus, error) {
	var ret Prometheus
	err := evaluator.EvaluateModule(ctx, source, &ret)
	return ret, err
}
