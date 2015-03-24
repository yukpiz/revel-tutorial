package controllers

import (
	"github.com/robfig/revel"
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	return c.Render()
}
