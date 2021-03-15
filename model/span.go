package model

// reference:https://opentracing.io/specification/
type Span struct {

	// An operation name
	Name string `json:"name"`

	// A start timestamp
	StartTimestamp uint64

	// A finish timestamp
	EndTimestamp uint64

	// A set of zero or more key:value Span Tags.
	// The keys must be strings. The values may be strings, bools, or numeric types.
	Tags []Tag `json:"tags"`

	// A set of zero or more Span Logs
	// each of which is itself a key:value map paired with a timestamp.
	// The keys must be strings, though the values may be of any type.
	// Not all OpenTracing implementations must support every value type.

	// A SpanContext (see below)
}
