package controllers

import (
	"github.com/revel/revel"
)

func (c User) Account() revel.Result {
	return c.Render()
}
