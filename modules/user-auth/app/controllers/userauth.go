package controllers

import (
	"fmt"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/modules/user-auth/app/models"
	"github.com/iassic/revel-modz/modules/user-auth/routes"
)

// interceptor to get
func (c *User) getUserLogin() revel.Result {
	// check if user is connected and logged in
	if user := c.connectedLogin(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	return nil
}

func (c *User) connectedLogin() (u *models.UserBase) {

	if userid, ok := c.RenderArgs["auth"].(int64); ok && userid > 0 {
		return c.getUserById(userid)
	}
	if useridstr, ok := c.Session["auth"]; ok {
		user := c.getUserByIdStr(useridstr)
		c.Args["user"] = user
		c.RenderArgs["user"] = user
		c.Args["auth"] = user.UserId
		c.RenderArgs["auth"] = user.UserId
		return user
	}
	return nil
}

func (c *App) Login() revel.Result {
	return c.Render()
}

func (c *App) LoginPost(email, password string) revel.Result {
	user := c.getUserByEmail(email)
	if user != nil {
		now := time.Now().UnixNano()
		auth := c.getUserAuth(user.UserId)
		volt := c.getUserVolt(user.UserId)
		if auth.LockExpiresAt > 0 {
			c.Flash.Out["email"] = email
			c.Flash.Error("Too many failed Logins, please contact support")
			return c.Redirect(routes.App.Login())
		}
		err := bcrypt.CompareHashAndPassword(auth.HashedPassword, []byte(password))
		if err == nil {
			auth.FailedLoginsCount = 0
			volt.LastLoginAt = now
			volt.LastActivityAt = now

			c.updateUserAuth(auth)
			c.updateUserVolt(volt)
			idstr := fmt.Sprint(user.UserId)
			c.Session["user"] = idstr
			c.Session["auth"] = idstr
			return c.Redirect(routes.User.Dashboard())
		} else {

			auth.FailedLoginsCount++
			volt.LastActivityAt = now
			if auth.FailedLoginsCount > 5 {
				// set lock on account
				auth.LockExpiresAt = 1000000
			}
			c.updateUserAuth(auth)
			c.updateUserVolt(volt)
		}
	}

	c.Flash.Out["email"] = email
	c.Flash.Error("Login failed")
	c.FlashParams()
	return c.Redirect(routes.App.Login())
}

func (c *App) Logout() revel.Result {
	user := c.RenderArgs["user"].(*models.UserBase)

	now := time.Now().UnixNano()
	volt := c.getUserVolt(user.UserId)
	volt.LastLogoutAt = now
	volt.LastActivityAt = now
	c.updateUserVolt(volt)

	for k := range c.Session {
		delete(c.Session, k)
	}
	for k := range c.Flash.Out {
		delete(c.Flash.Out, k)
	}

	return c.Redirect(routes.App.Index())
}
