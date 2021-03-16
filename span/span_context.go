package span

import (
	"sync"
)

type Baggage map[string]string

type ISpanContext interface {
	GetTraceId() uint64
	SetTraceId(id uint64)

	GetSpanId() uint64
	SetSpanId(id uint64)

	GetSpanParentId() uint64
	SetSpanParentId(spanId uint64)

	GetBaggageItem(key string) string
	SetBaggageItem(key, value string)
	GetAllBaggage() map[string]string
}

var _ ISpanContext = &Context{}

type Context struct {
	TraceId      uint64
	SpanId       uint64
	SpanParentId uint64

	// Baggage Items, which are just key:value pairs
	// that cross process boundaries
	Baggage Baggage
	mux     sync.RWMutex
}

func NewSpanContext() *Context {
	spanCtx := &Context{}
	return spanCtx
}

func (sc *Context) GetTraceId() uint64 {
	return sc.TraceId
}

func (sc *Context) SetTraceId(id uint64) {
	sc.TraceId = id
}

func (sc *Context) GetSpanId() uint64 {
	return sc.TraceId
}

func (sc *Context) SetSpanId(spanId uint64) {
	sc.SpanId = spanId
}

func (sc *Context) GetSpanParentId() uint64 {
	return sc.SpanParentId
}

func (sc *Context) SetSpanParentId(id uint64) {
	sc.SpanParentId = id
}

func (sc *Context) GetBaggageItem(key string) string {
	sc.mux.RLock()
	defer sc.mux.RUnlock()

	if sc.Baggage == nil {
		sc.Baggage = map[string]string{}
	}
	return sc.Baggage[key]
}

func (sc *Context) SetBaggageItem(key, value string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	if sc.Baggage == nil {
		sc.Baggage = map[string]string{}
	}
	sc.Baggage[key] = value
}

func (sc *Context) GetAllBaggage() map[string]string {
	sc.mux.RLock()
	defer sc.mux.RUnlock()

	temp := make(map[string]string)
	for k, v := range sc.Baggage {
		temp[k] = v
	}
	return temp
}
