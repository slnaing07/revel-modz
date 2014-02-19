package controllers

import (
	"fmt"

	"github.com/robfig/revel"

	"github.com\iassic\revel-modz\samples\/soe_smple/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) IndexPost(said string) revel.Result {
	c.Flash.Out["message"] = said
	return c.Redirect(routes.App.Result())
}

func (c App) Result() revel.Result {
	return c.Render()
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
