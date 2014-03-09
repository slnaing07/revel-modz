package controllers

import (
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"
)

type App struct {
	DbController
}

func (c App) RenderArgsFill() revel.Result {
	u := c.connected()
	if u != nil {
		c.RenderArgs["user_basic"] = u

		// look up role in RBAC module
		isAdmin := u.UserName == "admin@domain.com"
		if isAdmin {
			// set up things for an admin role
			c.Session["admin"] = "true"
		}
	}

	return nil
}

func (c App) connected() *user.UserBasic {
	if c.RenderArgs["user_basic"] != nil {
		return c.RenderArgs["user_basic"].(*user.UserBasic)
	}
	if username, ok := c.Session["user"]; ok {
		u := user.GetUserBasicByName(c.Txn, username)
		if u == nil {
			revel.ERROR.Println("user field in Session[] not found in DB")
			return nil
		}
		// revel.WARN.Printf("connected :: %+v", *u)
		return u
	}
	return nil
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Result() revel.Result {
	return c.Render()
}

// Defined in signup.go
// func (c App) Signup() revel.Result
// func (c App) SignupPost(usersignup *models.UserSignup) revel.Result
// func (c App) Register() revel.Result
// func (c App) RegisterPost(userregister *models.UserRegister) revel.Result
// func (c App) addNewUser(email, password string) (*user.UserBasic, error)

// Defined in maillist.go
// func (c App) Maillist() revel.Result
// func (c App) MaillistPost(usermaillist *models.UserMaillist) revel.Result
// func (c App) addNewMaillistUser(email, list string) (*maillist.MaillistUser, error)

// Defined in auth.go
// func (c App) Login() revel.Result
// func (c App) LoginPost(userlogin *models.UserLogin) revel.Result
// func (c App) Logout() revel.Result
