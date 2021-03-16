package span

import "time"

// reference:https://opentracing.io/specification/
type Span struct {

	// An operation name
	Name string `json:"name"`

	// A start timestamp
	StartTimestamp uint64

	// A finish timestamp - start timestamp
	Duration uint64

	// A set of zero or more key:value Span Tags.
	// The keys must be strings. The values may be strings, bools, or numeric types.
	Tags []Tag `json:"tags"`

	// A set of zero or more Span Logs
	// each of which is itself a key:value map paired with a timestamp.
	// The keys must be strings, though the values may be of any type.
	// Not all OpenTracing implementations must support every value type.
	Logs map[string]interface{}

	// A SpanContext (see below)
	ISpanContext
}

func (sp *Span) GetName() string {
	return sp.Name
}

func (sp *Span) SetName(name string) {
	sp.Name = name
}

func (sp *Span) GetSpanContext() ISpanContext {
	return sp.ISpanContext
}

func (sp *Span) SetSpanContext(ctx ISpanContext) {
	sp.ISpanContext = ctx
}

func (sp *Span) Finish() {

	finishTime := time.Now().UnixNano() / 1e3
	sp.Duration = uint64(finishTime) - sp.StartTimestamp
}
