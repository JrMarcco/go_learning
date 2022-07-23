package middleware

import (
	"go_learning/simple_web/framework"
	"net/http"
)

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatus(http.StatusInternalServerError).Json(err)
			}
		}()
		ctx.Next()
		return nil
	}
}