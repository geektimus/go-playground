package log

import (
	"context"
	"log"
)

// CorrelationID represents the id to track each request
type CorrelationID string

func (c CorrelationID) String() string {
	return string(c)
}

// Println prints a line duh!
func Println(ctx context.Context, msg string) {
	key := CorrelationID("correlation-id")
	id, ok := ctx.Value(key).(int64)
	if !ok {
		log.Println("Could not find request id in context")
		return
	}

	log.Printf("RequestID[%d]: %s", id, msg)
}
