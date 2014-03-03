package controllers

import (
	"github.com/revel/revel"
)

type Public struct {
	App
}

func (c Public) Documentation() revel.Result {
	return c.Render()
}

func (c Public) About() revel.Result {
	return c.Render()
}

func (c Public) Contact() revel.Result {
	return c.Render()
}
