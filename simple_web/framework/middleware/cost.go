package middleware

import (
	"go_learning/simple_web/framework"
	"log"
	"time"
)

func Cost() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		start := time.Now()

		ctx.Next()

		log.Printf("api uri: %v, cost: %v", ctx.Request().RequestURI, time.Now().Sub(start).Seconds())
		return nil
	}
}
