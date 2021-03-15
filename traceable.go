package tracing

type ITracer interface {
	// 把trace信息从ISpanContext中取出注入到ICarrier中
	Inject(ctx ISpanContext, carrier ICarrier) error

	// 从ICarrier中把trace信息取出放到ISpanContext中
	Extract(carrier ICarrier) (ISpanContext, error)
}
