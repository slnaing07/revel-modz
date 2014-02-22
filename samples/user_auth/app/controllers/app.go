package controllers

import (
	"fmt"

	"github.com/robfig/revel"

	"github.com/iassic/revel-modz/samples/user_auth/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Result(heading, message string) revel.Result {
	c.Flash.Out["heading"] = heading
	c.Flash.Out["message"] = message

	return c.Render()
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) SignupPost(email, password, confirmPassword string) revel.Result {
	c.Validation.Required(email)
	c.Validation.Required(password)

	c.Validation.Required(confirmPassword == user.Password).Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}
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
	c.Flash.Out["message"] = "You successfully registered."
	return c.Redirect(routes.App.Result())
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginPost(email, password) revel.Result {
	c.Validation.Required(email)
	c.Validation.Required(password)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Application.Register())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}
}
