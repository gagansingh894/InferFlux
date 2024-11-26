package engine

import "context"

type WorkloadEvaluator interface {
	StressAndBenchmark(ctx context.Context) error
}
