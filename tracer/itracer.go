package tracer

import (
	"github.com/think-next/tracing/carrier"
	"github.com/think-next/tracing/span"
)

const (
	ContextTracerKey = "tracer"
	ContextSpanKey   = "span"

	ContextTraceId       = "X-Trace-Id"
	ContextSpanId        = "X-Span-Id"
	ContextSpanParentId  = "X-Parent-Id"
	ContextFlag          = "X-Flag"
	ContextBaggagePrefix = "X-Baggage-"
)

// The Tracer interface creates Spans and understands
// how to Inject (serialize) and Extract (deserialize) them
// across process boundaries
type ITracer interface {
	Inject(ctx span.ISpanContext, carrier carrier.ICarrier) error
	Extract(carrier carrier.ICarrier) (span.ISpanContext, error)
}
