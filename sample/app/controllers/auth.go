package controllers

// TEMPLATE FILE

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
	UB, err := user.GetUserBasicByName(c.Txn, userlogin.Email)
	checkERROR(err)
	if UB != nil {
		found = true
	} else {
		c.Flash.Error("unknown user")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	// check for user in auth table
	passed, err := auth.Authenticate(c.Txn, UB.UserId, userlogin.Password)
	checkERROR(err)
	if !passed {
		c.Flash.Error("bad password")
	} else {
		valid = true
	}

	if found && valid {
		c.Session["user"] = UB.UserName
		c.RenderArgs["user_basic"] = UB

		// update visitor info in DB with UserId
		c.updateVisitorWithUserIdPanic()

		delete(c.Session, "v")
		delete(c.RenderArgs, "visitor")
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

	// update visitor info in DB with UserId
	c.updateVisitorWithUserIdPanic()
	// want to track them here, after they log out

	return c.Redirect(routes.App.Index())
}
