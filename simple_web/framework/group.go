package framework

type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	Use(middlewares ...ControllerHandler)

	Group(string) IGroup
}

type Group struct {
	core   *Core
	parent *Group
	prefix string

	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Group(url string) IGroup {
	childGroup := NewGroup(g.core, url)
	childGroup.parent = g
	return childGroup
}

func (g *Group) Get(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Get(url, all...)
}

func (g *Group) Post(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Post(url, all...)
}

func (g *Group) Put(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Put(url, all...)
}

func (g *Group) Delete(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Delete(url, all...)
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}
