package app

import (
	"context"
)

type App interface {
	Start(ctx context.Context)
	Shutdown(ctx context.Context)
}
