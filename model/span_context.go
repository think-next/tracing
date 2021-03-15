package model

import "sync"

type Baggage map[string]string

type SpanContext struct {
	traceID      uint64
	spanID       uint64
	parentSpanID uint64

	// Baggage Items, which are just key:value pairs that cross process boundaries
	baggage Baggage
	mux     sync.RWMutex
}

