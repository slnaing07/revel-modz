package controllers

import (
	"fmt"

	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
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

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) SignupPost(usersignup *models.UserSignup) revel.Result {
	usersignup.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	// check that this email is not in the DB already
	UB := user.GetUserBasicByName(c.Txn, usersignup.Email)
	if UB != nil {
		c.Validation.Error("Email already taken").Key("usersignup.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	UB, err := c.addNewUser(usersignup.Email, usersignup.Password)
	checkERROR(err)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = "Signup successful for " + usersignup.Email

	c.Session["user"] = UB.UserName
	c.RenderArgs["user_basic"] = UB
	return c.Redirect(routes.User.Result())

}

func (c App) Maillist() revel.Result {
	return c.Render()
}

func (c App) MaillistPost(usermaillist *models.UserMaillist) revel.Result {
	usermaillist.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Maillist())
	}

	// check that this email is not in the DB already
	UB := user.GetUserBasicByName(c.Txn, usermaillist.Email)
	if UB != nil {
		c.Validation.Error("Email already taken").Key("usermaillist.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	_, err := c.addNewMaillistUser(usermaillist.Email)
	checkERROR(err)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = usermaillist.Email + " is now subscribed to the mailing list."

	return c.Redirect(routes.App.Result())

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
		found = true
	} else {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	// check for user in auth table
	P := user.UserPass{UB.UserId, email, password}
	U, err := auth.Authenticate(c.Txn, &P)
	if err != nil || U == nil {
		revel.WARN.Println(err)
	} else {
		valid = true
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
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) addNewUser(email, password string) (*user.UserBasic, error) {

	// uuid := get random number (that isn't used already)
	uuid := user.GenerateNewUserId(c.Txn)
	UB := &user.UserBasic{
		UserId:   uuid,
		UserName: email,
	}
	UP := &user.UserPass{UB.UserId, email, password}

	// add user to tables
	// TODO do something more with the errosr
	err := user.AddUserBasic(TestDB, UB)
	checkERROR(err)

	_, err = auth.AddUserAuth(TestDB, UP)
	checkERROR(err)

	return UB, nil
}

func (c App) addNewMaillistUser(email string) (*maillist.MaillistUser, error) {

	// uuid := get random number (that isn't used already)
	uuid := user.GenerateNewUserId(c.Txn)
	UB := &user.UserBasic{
		UserId:   uuid,
		UserName: email,
	}

	err := user.AddUserBasic(TestDB, UB)
	checkERROR(err)

	MA, err := maillist.AddMaillistUser(TestDB, uuid, email)
	checkERROR(err)

	return MA, nil
}
