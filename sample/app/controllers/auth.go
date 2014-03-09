package controllers

import (
	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
)

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginPost(userlogin *models.UserLogin) revel.Result {
	userlogin.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	var found, valid bool

	// check for user in basic table
	UB := user.GetUserBasicByName(c.Txn, userlogin.Email)
	if UB != nil {
		found = true
	} else {
		println("FLASH ERROR USER")
		c.Flash.Error("unknown user")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	// check for user in auth table
	P := user.UserPass{UB.UserId, userlogin.Email, userlogin.Password}
	U, err := auth.Authenticate(c.Txn, &P)
	if err != nil || U == nil {
		println("FLASH ERROR PASSWORD")
		c.Flash.Error("bad password")
	} else {
		valid = true
	}

	if found && valid {
		c.Session["user"] = UB.UserName
		c.RenderArgs["user_basic"] = UB
		return c.Redirect(routes.User.Result())

	} else {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}
