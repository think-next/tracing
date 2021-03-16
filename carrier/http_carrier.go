package carrier

import (
	"github.com/think-next/tracing/span"
	"github.com/think-next/tracing/tracer"
	"net/http"
	"strconv"
	"strings"
)

// force HttpCarrier implement Carrier interface
var _ ICarrier = &HttpCarrier{}

type HttpCarrier struct {
	header http.Header
}

func NewHttpCarrier(header http.Header) *HttpCarrier {
	hc := &HttpCarrier{}
	hc.header = header
	return hc
}

func (hc *HttpCarrier) Inject(ctx span.ISpanContext) {

	hc.header.Set(tracer.ContextTraceId, strconv.FormatUint(ctx.GetTraceId(), 16))
	hc.header.Set(tracer.ContextSpanId, strconv.FormatUint(0, 16))
	hc.header.Set(tracer.ContextSpanParentId, strconv.FormatUint(ctx.GetSpanParentId(), 16))

	for k, v := range ctx.GetAllBaggage() {
		hc.header.Add(tracer.ContextBaggagePrefix+k, v)
	}
}

func (hc *HttpCarrier) Extract(ctx span.ISpanContext) {
	for k := range hc.header {
		switch k {
		case tracer.ContextTraceId:
			traceId, _ := strconv.ParseUint(hc.header.Get(k), 16, 64)
			ctx.SetTraceId(traceId)
		case tracer.ContextSpanId:
			spanId, _ := strconv.ParseUint(hc.header.Get(k), 16, 64)
			ctx.SetSpanId(spanId)
		case tracer.ContextSpanParentId:
			parentId, _ := strconv.ParseUint(hc.header.Get(k), 16, 64)
			ctx.SetSpanParentId(parentId)
		default:
			if strings.HasPrefix(k, tracer.ContextBaggagePrefix) {
				ctx.SetBaggageItem(k, hc.header.Get(k))
			}
		}
	}
}
