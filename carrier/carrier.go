package carrier

import "github.com/think-next/tracing/span"

type ICarrier interface {

	// Extract a Context from a carrier
	Extract(span.ISpanContext)

	// Inject a Context into a carrier
	Inject(span.ISpanContext)
}


