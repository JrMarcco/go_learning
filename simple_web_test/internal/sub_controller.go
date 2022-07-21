package internal

import "go_learning/simple_web/framework"

func SubjectAddController(c *framework.Context) error {
	c.OkJson("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.OkJson("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.OkJson("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.OkJson("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.OkJson("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.OkJson("ok, SubjectNameController")
	return nil
}
