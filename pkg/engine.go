package pkg

import "context"

type LoadTester interface {
	Start(ctx context.Context) error
}
