package controllers

import (
	"github.com/robfig/revel"
)

type App2 struct {
	*revel.Controller
}

func (c App2) Index() revel.Result {
	return c.Render()
}
