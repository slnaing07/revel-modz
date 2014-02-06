package controllers

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
)

type AuthController struct {
	revel.Controller
}

func (c *AuthController) GetUserLogin() revel.Result {
	if useridstr, ok := c.Session["auth-userid"]; ok {
		user := c.getUserByIdStr(useridstr)
		c.Args["user"] = user
		c.RenderArgs["user"] = user
		c.Args["auth"] = user.UserId
		c.RenderArgs["auth"] = user.UserId
		return user
	}
}
