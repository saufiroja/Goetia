package main

import (
	"context"
	"github.com/saufiroja/cqrs/internal/app"
)

func main() {
	start := app.NewAppFactor()
	start.StartApp(context.Background())
}
