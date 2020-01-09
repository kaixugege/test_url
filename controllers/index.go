package controllers

type IndexController struct {
	BaseConroller
}

// @router / [get]
func (c *IndexController) Index() {
	c.TplName = "index.html"
}

// @router /about [get]
func (c *IndexController) IndexAbout() {
	c.TplName = "about.html"
}

// @router /message [get]
func (c *IndexController) IndexMessage() {
	c.Abort("404")
	c.TplName = "message.html"
}

// @router /user [get]
func (c *IndexController) IndexUser() {
	c.TplName = "user.html"
}
