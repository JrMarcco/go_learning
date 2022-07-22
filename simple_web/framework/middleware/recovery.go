package middleware

import "go_learning/simple_web/framework"

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				ctx.ErrJson(err)
			}
		}()
		ctx.Next()
		return nil
	}
}
