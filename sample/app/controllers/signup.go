package controllers

import (
	"github.com/iassic/revel-modz/modules/auth"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
)

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

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) RegisterPost(userregister *models.UserRegister) revel.Result {
	userregister.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Maillist())
	}

	// check that this email is not in the DB already
	UB := user.GetUserBasicByName(c.Txn, userregister.Email)
	if UB != nil {
		c.Validation.Error("Email already taken").Key("userregister.Email")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	var err error
	UB, err = c.addNewUser(userregister.Email, userregister.Password)
	checkERROR(err)

	// TODO  which mailing lists did they check off?
	// ALSO  user Basic will be added twice if this current call is made
	// _, err = c.addNewMaillistUser(userregister.Email)
	// checkERROR(err)

	// TODO add profile DB insert

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = userregister.Email + " is now subscribed to the mailing list."

	return c.Redirect(routes.App.Result())

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
