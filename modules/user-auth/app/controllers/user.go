package controllers

import (
	"fmt"
	// "time"

	// "code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"
	"iassic-demo/app/models"
	// "iassic-demo/app/routes"
)

type User struct {
	App
}

func (c *User) Index() revel.Result {
	return c.Render()
}

func (c *User) Profile() revel.Result {
	return c.Render()
}

func (c *User) Account() revel.Result {
	return c.Render()
}

func (c *User) Dashboard() revel.Result {
	return c.Render()
}

func (c *App) getUserById(userid int64) *models.UserBase {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserBase{}, fmt.Sprintf("SELECT * FROM userbase WHERE userid = '%d'", userid))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserBase)
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

func (c *App) getUserByName(username string) *models.UserBase {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserBase{}, fmt.Sprintf("SELECT * FROM userbase WHERE username = '%s'", username))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserBase)
}

func (c *App) getUserByEmail(email string) *models.UserBase {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserBase{}, fmt.Sprintf("SELECT * FROM userbase WHERE email = '%s'", email))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserBase)
}

func (c *App) checkUserExists(username, email string) (name_exists, email_exists bool) {
	u := c.getUserByName(username)
	if u != nil {
		name_exists = true
	}

	e := c.getUserByEmail(email)
	if e != nil {
		email_exists = true
	}

	return
}

func (c *App) getUserAuth(userid int64) *models.UserAuth {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserAuth{}, fmt.Sprintf("SELECT * FROM userauth WHERE userid = '%d'", userid))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserAuth)

}

func (c *App) getUserVolt(userid int64) *models.UserVolatile {
	// TODO  SQL prepared statement
	users, err := c.Txn.Select(models.UserVolatile{}, fmt.Sprintf("SELECT * FROM uservolatile WHERE userid = '%d'", userid))
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.UserVolatile)

}

func (c *App) updateUserBase(user *models.UserBase) error {
	_, errU := c.Txn.Update(user)
	if errU != nil {
		return errU
	}
	return nil
}

// XXX   this seems to update all fields, probably a bad idea
//       should write individual functions to handle updating each field separately, or only those that should be changing
func (c *App) updateUserAuth(auth *models.UserAuth) error {
	_, errU := c.Txn.Update(auth)
	if errU != nil {
		return errU
	}
	return nil
}

func (c *App) updateUserVolt(volt *models.UserVolatile) error {
	_, errU := c.Txn.Update(volt)
	if errU != nil {
		return errU
	}
	return nil
}
