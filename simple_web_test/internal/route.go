package internal

import (
	"go_learning/simple_web/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("/user/login", UserController)

	subApi := core.Group("/sub")
	{
		subApi.Get("/:id", SubjectGetController)
		subApi.Put("/:id", SubjectUpdateController)
		subApi.Delete("/:id", SubjectDelController)
		subApi.Get("/list/all", SubjectListController)

		subInnerApi := subApi.Group("/info")
		{
			subInnerApi.Get("/name", SubjectNameController)
		}
	}
}
