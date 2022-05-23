package filter

import (
	"fmt"
	"go_learning/simple-web/pkg/context"
	"time"
)

type HttpFiler func(ctx *context.HttpContext)

type HttpFilterBuilder func(next HttpFiler) HttpFiler

var _ HttpFilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next HttpFiler) HttpFiler {
	return func(ctx *context.HttpContext) {
		startTime := time.Now().UnixNano()

		method := ctx.Req.Method
		path := ctx.Req.URL.Path
		next(ctx)
		fmt.Printf("### [%s] [%s] take %d ns ###\n", method, path, time.Now().UnixNano()-startTime)
	}
}
