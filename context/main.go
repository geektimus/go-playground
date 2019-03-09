package main

import (
	"context"

	"github.com/geektimus/getting-started/context/log"
)

func main() {
	key := log.CorrelationID("correlation-id")
	ctx := context.WithValue(context.Background(), key, int64(22222222))
	log.Println(ctx, "Hello World!")
}
