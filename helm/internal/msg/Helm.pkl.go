// Code generated from Pkl module `pkl.helm.helm`. DO NOT EDIT.
package msg

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Helm struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Helm
func LoadFromPath(ctx context.Context, path string) (ret Helm, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Helm
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Helm, error) {
	var ret Helm
	err := evaluator.EvaluateModule(ctx, source, &ret)
	return ret, err
}
