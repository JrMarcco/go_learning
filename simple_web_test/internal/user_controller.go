package internal

import "go_learning/simple_web/framework"

func UserController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, UserController")
	return nil
}
