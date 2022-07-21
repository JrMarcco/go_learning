package internal

import "go_learning/simple_web/framework"

func UserController(ctx *framework.Context) error {
	ctx.OkJson("ok, UserController")
	return nil
}
