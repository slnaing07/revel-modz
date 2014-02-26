package controllers

import (
	// "fmt"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"
	"iassic-demo/app/models"
	"iassic-demo/app/routes"
)

func (c *App) Signup() revel.Result {
	return c.Render()
}

func (c *App) SaveUser(user models.UserBase, auth models.UserAuth, password, verifyPassword string) revel.Result {
	user.UserName = user.Email[:strings.Index(user.Email, "@")]
	_, e_exists := c.checkUserExists(user.UserName, user.Email)

	if e_exists {
		msg := "Sorry, "
		if e_exists {
			c.Flash.Data["user.Email"] = ""
			msg += "an account is already registered with that email address."
		}
		c.Flash.Error(msg)
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	auth.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)

	err := c.insertNewUser(user, auth)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.UserName
	c.Flash.Success("Welcome, " + user.UserName)
	return c.Redirect(routes.App.Index())
}

func (c *App) insertNewUser(user models.UserBase, auth models.UserAuth) error {
	now := time.Now().UnixNano()

	errB := c.Txn.Insert(&user)
	if errB != nil {
		return errB
	}

	auth.UserId = user.UserId
	errA := c.Txn.Insert(&auth)
	if errA != nil {
		return errA
	}

	var volt models.UserVolatile
	volt.UserId = user.UserId
	volt.CreatedAt = now
	volt.LastLoginAt = now
	volt.LastActivityAt = now

	errV := c.Txn.Insert(&volt)
	if errV != nil {
		return errV
	}

	return nil
}
