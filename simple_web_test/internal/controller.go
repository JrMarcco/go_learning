package internal

import (
	"context"
	"fmt"
	"go_learning/simple_web/framework"
	"log"
	"net/http"
	"time"
)

func FooController(ctx *framework.Context) error {
	done := make(chan struct{}, 1)
	panicChan := make(chan any, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 5*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.SetOkStatus().Json("ok")

		done <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Default().Println(p)
		ctx.SetStatus(http.StatusInternalServerError).Json("panic")
	case <-done:
		fmt.Println("done")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.SetStatus(http.StatusInternalServerError).Json("panic")
		ctx.SetTimeout()
	}
	return nil
}
