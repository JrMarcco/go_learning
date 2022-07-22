package middleware

import (
	"context"
	"fmt"
	"go_learning/simple_web/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(ctx *framework.Context) error {

		done := make(chan struct{}, 1)
		panicChan := make(chan any, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			ctx.Next()

			done <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Default().Println(p)
			ctx.ErrJson("panic")
		case <-done:
			fmt.Println("done")
		case <-durationCtx.Done():
			ctx.ErrJson("timeout")
			ctx.SetTimeout()
		}

		return nil
	}
}
