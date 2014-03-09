package controllers

import (
	"github.com/iassic/revel-modz/modules/maillist"
	"github.com/iassic/revel-modz/modules/user"
	"github.com/revel/revel"

	"github.com/iassic/revel-modz/sample/app/models"
	"github.com/iassic/revel-modz/sample/app/routes"
)

// Anyone functions
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

	_, err := c.addNewMaillistUser(usermaillist.Email, "MaillistPost()")
	checkERROR(err)

	c.Flash.Out["heading"] = "Thanks for Joining!"
	c.Flash.Out["message"] = usermaillist.Email + " is now subscribed to the mailing list."

	return c.Redirect(routes.App.Result())

}

// Admin functions
func (c Admin) MaillistView() revel.Result {
	maillist_users, err := maillist.GetAllUsers(c.Txn)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Render()
	}
	return c.Render(maillist_users)
}

func (c Admin) MaillistCompose() revel.Result {
	return c.Render()
}

// helper functions

func (c App) addNewMaillistUser(email, list string) (*maillist.MaillistUser, error) {

	// uuid := get random number (that isn't used already)
	uuid := user.GenerateNewUserId(c.Txn)
	UB := &user.UserBasic{
		UserId:   uuid,
		UserName: email,
	}

	err := user.AddUserBasic(TestDB, UB)
	checkERROR(err)

	MA, err := maillist.AddUser(TestDB, uuid, email, list)
	checkERROR(err)

	return MA, nil
}
