package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	Group(string) IGroup
}

type Group struct {
	core   *Core
	parent *Group
	prefix string
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

func (g *Group) GetAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.GetAbsolutePrefix() + g.prefix
}

func (g *Group) Get(url string, handler ControllerHandler) {
	url = g.GetAbsolutePrefix() + url
	g.core.Get(url, handler)
}

func (g *Group) Post(url string, handler ControllerHandler) {
	url = g.GetAbsolutePrefix() + url
	g.core.Post(url, handler)
}

func (g *Group) Put(url string, handler ControllerHandler) {
	url = g.GetAbsolutePrefix() + url
	g.core.Put(url, handler)
}

func (g *Group) Delete(url string, handler ControllerHandler) {
	url = g.GetAbsolutePrefix() + url
	g.core.Delete(url, handler)
}
