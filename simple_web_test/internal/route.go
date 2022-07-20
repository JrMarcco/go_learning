package internal

import (
	"go_learning/simple_web/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
