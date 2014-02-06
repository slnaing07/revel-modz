package controllers

import (
	// "github.com/coopernurse/gorp"
	"github.com/robfig/revel"
)

type AuthController struct {
	revel.Controller
	Txn *gorp.Transaction
}

func (c *AuthController) GetUserLogin() revel.Result {
	if useridstr, ok := c.Session["auth-userid"]; ok {
		user := c.getUserByIdStr(useridstr)
		c.Args["auth-userid"] = user.UserId
		c.RenderArgs["auth-userid"] = user.UserId
		return user
	}
}

func (c *App) getUserByIdStr(userid string) *models.UserBase {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserBase{}, fmt.Sprintf("SELECT * FROM userbase WHERE userid = '%s'", userid))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserBase)
}
