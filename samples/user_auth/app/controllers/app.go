package controllers

import (
	"fmt"

	"github.com/robfig/revel"

	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/iassic/revel-modz/samples/user_auth/app/routes"
)

type App struct {
	DbController
}

func (c App) RenderArgsFill() revel.Result {
	if u := c.connected(); u != nil {
		c.RenderArgs["user_basic"] = u
	}
	return nil
}

func (c App) connected() *user.UserBasic {
	if c.RenderArgs["user_basic"] != nil {
		return c.RenderArgs["user_basic"].(*user.UserBasic)
	}
	if username, ok := c.Session["user_basic"]; ok {
		u := user.GetUserBasicByName(c.Txn, username)
		if u == nil {
			revel.WARN.Println("user_basic field in Session[] not found in DB")
			return nil
		}
		return u
	}
	return nil
}

func (c App) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.User.Index())
	}
	return c.Render()
}

func (c App) Result() revel.Result {
	return c.Render()
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) SignupPost(email, password, confirmPassword string) revel.Result {
	c.Validation.Required(email).Message("Email required")
	c.Validation.Required(password).Message("Password required")

	c.Validation.Required(password == confirmPassword).Message("Passwords do not match")

	// user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}
	// Check email is unique
	UB := user.GetUserBasicByName(c.Txn, email)
	if UB != nil {
		revel.WARN.Printf("Duplicate Email found %+v\n", UB)
		c.Flash.Out["message"] = "Email: " + email + " already used"
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = "Signup successful for " + email

	c.Session["user"] = UB.UserName
	c.RenderArgs["user_basic"] = UB
	return c.Redirect(routes.User.Result())

}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) RegisterPost(fname, mi, lname, email, dob, sex, address, city, state, zipcode, phonenumber string) revel.Result {
	fmt.Println("fname", fname)
	fmt.Println("mi", mi)
	fmt.Println("lname", lname)
	fmt.Println("email", email)
	fmt.Println("dob", dob)
	fmt.Println("sex", sex)
	fmt.Println("address", address)
	fmt.Println("city", city)
	fmt.Println("state", state)
	fmt.Println("zipcode", zipcode)
	fmt.Println("phonenumber", phonenumber)
	c.Flash.Out["heading"] = "RegisterPost"
	c.Flash.Out["message"] = "You sorta-successfully fake-registered."
	return c.Redirect(routes.App.Result())
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginPost(email, password string) revel.Result {
	var found, valid bool

	c.Validation.Required(email).Message("Email required")
	c.Validation.Required(password).Message("Password required")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	// check for user in basic table
	UB := user.GetUserBasicByName(c.Txn, email)
	if UB != nil {
		revel.WARN.Printf("%+v\n", UB)
		found = true
	} else {
		revel.WARN.Println("User not found: ", email)
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	// check for user in auth table
	P := user.UserPass{UB.UserId, email, password}
	U, err := auth.Authenticate(c.Txn, &P)
	if err != nil {
		revel.WARN.Println(err)
	} else {
		valid = true
		revel.INFO.Println(U)
	}

	if found && valid {
		c.Flash.Out["heading"] = "LOGIN PASS"
		c.Flash.Out["message"] = "Login successful for " + email

		c.Session["user"] = UB.UserName
		c.RenderArgs["user_basic"] = UB
		return c.Redirect(routes.User.Result())

	} else {
		c.Flash.Out["heading"] = "LOGIN FAIL"
		c.Flash.Out["message"] = "Login failed for " + email
		c.Validation.Keep()
		c.FlashParams()
		c.Redirect(routes.App.Login())
	}
	return c.Redirect(routes.App.Result())
}

func (c App) Logout() revel.Result {
	fmt.Printf("Deleting session keys...\n")
	for k := range c.Session {
		fmt.Printf("Deleting Session[%s]: '%s'\n", k, c.Session[k])
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}
