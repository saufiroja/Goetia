package main

import (
	_ "github.com/lib/pq"
	"github.com/saufiroja/cqrs/internal"
)

func main() {
	internal.Start()
}
